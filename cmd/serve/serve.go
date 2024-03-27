// Copyright 2020 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package serve

import (
	"context"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/avast/retry-go"
	clusterv3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpointv3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	tlsv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	runtimev3 "github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	ratelimitv3 "github.com/envoyproxy/go-control-plane/ratelimit/config/ratelimit/v3"
	"github.com/mitchellh/cli"
	"github.com/stephenafamo/orchestra"

	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"github.com/envoyproxy/go-control-plane/pkg/server/v3"
	"github.com/envoyproxy/go-control-plane/pkg/test/v3"
	"github.com/protoconf/libprotoconf"
	"github.com/protoconf/protoconf-xds/src/protoconfxds/v1"
	"github.com/protoconf/protoconf-xds/xds"
	protoconf_agent "github.com/protoconf/protoconf/agent/api/proto/v1"
	"github.com/smintz/keygroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Command struct {
	logger  *xds.Logger
	config  *protoconfxds.ProtoconfEnvoyConfig
	flagset *flag.FlagSet
}

func (c *Command) Synopsis() string {
	return "Run the protoconf-xds server"
}

func (c *Command) Help() string {
	buf := &strings.Builder{}
	c.flagset.SetOutput(buf)
	c.flagset.Usage()
	return fmt.Sprintf("%s\n\n%v", c.Synopsis(), buf.String())
}

func NewCommand() (cli.Command, error) {
	config := &protoconfxds.ProtoconfEnvoyConfig{
		ProtoconfAgentAddr: "localhost:4300",
		Prefix:             "example",
		Port:               18000,
	}
	lpc := libprotoconf.NewConfig(config)
	lpc.SetEnvKeyPrefix("PROTOCONF_XDS")
	err := lpc.Environment()
	if err != nil {
		return nil, err
	}
	command := &Command{
		logger:  &xds.Logger{},
		flagset: lpc.DefaultFlagSet(),
		config:  config,
	}
	return command, nil

}

func (c *Command) Run(args []string) int {
	err := c.flagset.Parse(args)
	if err != nil {
		c.flagset.Usage()
		return 1
	}
	c.logger.Debug = c.config.Debug

	// Create a cache
	xdsCache := cache.NewSnapshotCache(false, cache.IDHash{}, c.logger)

	conn, err := grpc.Dial(
		"localhost:4300",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		c.logger.Errorf("failed to connect to agent", "error", err)
		return 1
	}
	defer conn.Close()
	stub := protoconf_agent.NewProtoconfServiceClient(conn)
	watcher := func(ctx context.Context, key string) {
		retry.Do(func() error {
			stream, err := stub.SubscribeForConfig(ctx, &protoconf_agent.ConfigSubscriptionRequest{
				Path: filepath.Join(c.config.Prefix, key),
			})
			if err != nil {
				c.logger.Errorf("error getting stream %v", err)
			}
			xdsRecord := &protoconfxds.XDSSnapshot{}
			for {
				update, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if status.Code(err) == codes.Canceled {
					c.logger.Infof("stopping %s", key)
					break
				}
				if err != nil {
					c.logger.Errorf("%v", err)
					return err
				}
				err = update.Value.UnmarshalTo(xdsRecord)
				if err != nil {
					c.logger.Errorf("%v", err)
					return err
				}
				version := fmt.Sprintf("%x", md5.Sum([]byte(update.String())))
				c.logger.Debugf("version: %s : %v", version, xdsRecord)
				snapshot, err := cache.NewSnapshot(version, map[resource.Type][]types.Resource{
					resource.ClusterType:         makeClusters(xdsRecord.Clusters),
					resource.RouteType:           makeRoutes(xdsRecord.Routes),
					resource.ListenerType:        makeListeners(xdsRecord.Listeners),
					resource.EndpointType:        makeEndpoints(xdsRecord.Endpoints),
					resource.ScopedRouteType:     makeScopedRoutes(xdsRecord.ScopedTypes),
					resource.VirtualHostType:     makeVirtualHosts(xdsRecord.VirtualHosts),
					resource.SecretType:          makeSecrets(xdsRecord.Secrets),
					resource.ExtensionConfigType: makeExtensionConfigs(xdsRecord.ExtentionConfigs),
					resource.RuntimeType:         makeRuntimes(xdsRecord.Runtimes),
					resource.RateLimitConfigType: makeRatelimits(xdsRecord.Ratelimits),
				})
				if err != nil {
					c.logger.Errorf("%v", err)
					return err
				}
				err = xdsCache.SetSnapshot(ctx, key, snapshot)
				if err != nil {
					c.logger.Errorf("%v", err)
				}
			}
			return nil
		})
	}
	kg := keygroup.NewKeyGroup(watcher)
	kg.Update(c.config.NodeId)

	// Run the xDS server
	ctx := context.Background()
	cb := &test.Callbacks{Debug: c.logger.Debug}
	srv := server.NewServer(ctx, xdsCache, cb)
	err = orchestra.PlayUntilSignal(orchestra.PlayerFunc(xds.GeneratePlayer(srv, uint(c.config.Port))), syscall.SIGINT, syscall.SIGTERM)
	if err != nil {
		c.logger.Errorf("grpc server return error: %v", err)
		return 1
	}
	return 0
}

func makeClusters(input []*clusterv3.Cluster) (ret []types.Resource) {
	for _, item := range input {
		ret = append(ret, item)
	}
	return ret
}
func makeListeners(input []*listenerv3.Listener) (ret []types.Resource) {
	for _, item := range input {
		ret = append(ret, item)
	}
	return ret
}
func makeRoutes(input []*routev3.RouteConfiguration) (ret []types.Resource) {
	for _, item := range input {
		ret = append(ret, item)
	}
	return ret
}
func makeEndpoints(input []*endpointv3.ClusterLoadAssignment) (ret []types.Resource) {
	for _, item := range input {
		ret = append(ret, item)
	}
	return ret
}
func makeScopedRoutes(input []*routev3.ScopedRouteConfiguration) (ret []types.Resource) {
	for _, item := range input {
		ret = append(ret, item)
	}
	return ret
}
func makeVirtualHosts(input []*routev3.VirtualHost) (ret []types.Resource) {
	for _, item := range input {
		ret = append(ret, item)
	}
	return ret
}
func makeSecrets(input []*tlsv3.Secret) (ret []types.Resource) {
	for _, item := range input {
		ret = append(ret, item)
	}
	return ret
}
func makeExtensionConfigs(input []*corev3.TypedExtensionConfig) (ret []types.Resource) {
	for _, item := range input {
		ret = append(ret, item)
	}
	return ret
}
func makeRuntimes(input []*runtimev3.Runtime) (ret []types.Resource) {
	for _, item := range input {
		ret = append(ret, item)
	}
	return ret
}

func makeRatelimits(input []*ratelimitv3.RateLimitConfig) (ret []types.Resource) {
	for _, item := range input {
		ret = append(ret, item)
	}
	return ret
}
