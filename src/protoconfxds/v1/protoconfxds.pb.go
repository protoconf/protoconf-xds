// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: src/protoconfxds/v1/protoconfxds.proto

package protoconfxds

import (
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	v35 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	v33 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	v32 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	v37 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/v3"
	v34 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	v36 "github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProtoconfEnvoyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProtoconfAgentAddr string   `protobuf:"bytes,1,opt,name=protoconf_agent_addr,json=protoconfAgentAddr,proto3" json:"protoconf_agent_addr,omitempty"`
	Prefix             string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	NodeId             []string `protobuf:"bytes,3,rep,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	Port               uint32   `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Debug              bool     `protobuf:"varint,5,opt,name=debug,proto3" json:"debug,omitempty"`
}

func (x *ProtoconfEnvoyConfig) Reset() {
	*x = ProtoconfEnvoyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protoconfxds_v1_protoconfxds_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProtoconfEnvoyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoconfEnvoyConfig) ProtoMessage() {}

func (x *ProtoconfEnvoyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_src_protoconfxds_v1_protoconfxds_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoconfEnvoyConfig.ProtoReflect.Descriptor instead.
func (*ProtoconfEnvoyConfig) Descriptor() ([]byte, []int) {
	return file_src_protoconfxds_v1_protoconfxds_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoconfEnvoyConfig) GetProtoconfAgentAddr() string {
	if x != nil {
		return x.ProtoconfAgentAddr
	}
	return ""
}

func (x *ProtoconfEnvoyConfig) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *ProtoconfEnvoyConfig) GetNodeId() []string {
	if x != nil {
		return x.NodeId
	}
	return nil
}

func (x *ProtoconfEnvoyConfig) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ProtoconfEnvoyConfig) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

type XDSSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoints        []*v3.ClusterLoadAssignment     `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	Clusters         []*v31.Cluster                  `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	Routes           []*v32.RouteConfiguration       `protobuf:"bytes,3,rep,name=routes,proto3" json:"routes,omitempty"`
	ScopedTypes      []*v32.ScopedRouteConfiguration `protobuf:"bytes,4,rep,name=scoped_types,json=scopedTypes,proto3" json:"scoped_types,omitempty"`
	VirtualHosts     []*v32.VirtualHost              `protobuf:"bytes,5,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	Listeners        []*v33.Listener                 `protobuf:"bytes,6,rep,name=listeners,proto3" json:"listeners,omitempty"`
	Secrets          []*v34.Secret                   `protobuf:"bytes,7,rep,name=secrets,proto3" json:"secrets,omitempty"`
	ExtentionConfigs []*v35.TypedExtensionConfig     `protobuf:"bytes,8,rep,name=extention_configs,json=extentionConfigs,proto3" json:"extention_configs,omitempty"`
	Runtimes         []*v36.Runtime                  `protobuf:"bytes,9,rep,name=runtimes,proto3" json:"runtimes,omitempty"`
	ThriftRoutes     []*v37.RouteConfiguration       `protobuf:"bytes,10,rep,name=thrift_routes,json=thriftRoutes,proto3" json:"thrift_routes,omitempty"`
}

func (x *XDSSnapshot) Reset() {
	*x = XDSSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_protoconfxds_v1_protoconfxds_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XDSSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XDSSnapshot) ProtoMessage() {}

func (x *XDSSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_src_protoconfxds_v1_protoconfxds_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XDSSnapshot.ProtoReflect.Descriptor instead.
func (*XDSSnapshot) Descriptor() ([]byte, []int) {
	return file_src_protoconfxds_v1_protoconfxds_proto_rawDescGZIP(), []int{1}
}

func (x *XDSSnapshot) GetEndpoints() []*v3.ClusterLoadAssignment {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *XDSSnapshot) GetClusters() []*v31.Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *XDSSnapshot) GetRoutes() []*v32.RouteConfiguration {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *XDSSnapshot) GetScopedTypes() []*v32.ScopedRouteConfiguration {
	if x != nil {
		return x.ScopedTypes
	}
	return nil
}

func (x *XDSSnapshot) GetVirtualHosts() []*v32.VirtualHost {
	if x != nil {
		return x.VirtualHosts
	}
	return nil
}

func (x *XDSSnapshot) GetListeners() []*v33.Listener {
	if x != nil {
		return x.Listeners
	}
	return nil
}

func (x *XDSSnapshot) GetSecrets() []*v34.Secret {
	if x != nil {
		return x.Secrets
	}
	return nil
}

func (x *XDSSnapshot) GetExtentionConfigs() []*v35.TypedExtensionConfig {
	if x != nil {
		return x.ExtentionConfigs
	}
	return nil
}

func (x *XDSSnapshot) GetRuntimes() []*v36.Runtime {
	if x != nil {
		return x.Runtimes
	}
	return nil
}

func (x *XDSSnapshot) GetThriftRoutes() []*v37.RouteConfiguration {
	if x != nil {
		return x.ThriftRoutes
	}
	return nil
}

var File_src_protoconfxds_v1_protoconfxds_proto protoreflect.FileDescriptor

var file_src_protoconfxds_v1_protoconfxds_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x78,
	0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x78,
	0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6e, 0x66, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f,
	0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x64, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x75, 0x6e,
	0x74, 0x69, 0x6d, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x74, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa3, 0x01, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x45, 0x6e,
	0x76, 0x6f, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6e, 0x66, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x22, 0x8c, 0x06, 0x0a, 0x0b, 0x58, 0x44, 0x53, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x4d, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x12, 0x41, 0x0a, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x52, 0x0a, 0x0c, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x73,
	0x63, 0x6f, 0x70, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x47, 0x0a, 0x0d, 0x76, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x0c, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x40, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76,
	0x33, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x52, 0x09, 0x6c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x4b, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65,
	0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x2e, 0x74, 0x6c, 0x73, 0x2e,
	0x76, 0x33, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x52, 0x07, 0x73, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x73, 0x12, 0x57, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x3d, 0x0a, 0x08, 0x72,
	0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x72, 0x75,
	0x6e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x52, 0x08, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x69, 0x0a, 0x0d, 0x74, 0x68,
	0x72, 0x69, 0x66, 0x74, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x44, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x68, 0x72, 0x69, 0x66, 0x74, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x78, 0x64, 0x73, 0x2f, 0x76,
	0x31, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x78, 0x64, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_protoconfxds_v1_protoconfxds_proto_rawDescOnce sync.Once
	file_src_protoconfxds_v1_protoconfxds_proto_rawDescData = file_src_protoconfxds_v1_protoconfxds_proto_rawDesc
)

func file_src_protoconfxds_v1_protoconfxds_proto_rawDescGZIP() []byte {
	file_src_protoconfxds_v1_protoconfxds_proto_rawDescOnce.Do(func() {
		file_src_protoconfxds_v1_protoconfxds_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_protoconfxds_v1_protoconfxds_proto_rawDescData)
	})
	return file_src_protoconfxds_v1_protoconfxds_proto_rawDescData
}

var file_src_protoconfxds_v1_protoconfxds_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_protoconfxds_v1_protoconfxds_proto_goTypes = []interface{}{
	(*ProtoconfEnvoyConfig)(nil),         // 0: protoconfenvoy.v1.ProtoconfEnvoyConfig
	(*XDSSnapshot)(nil),                  // 1: protoconfenvoy.v1.XDSSnapshot
	(*v3.ClusterLoadAssignment)(nil),     // 2: envoy.config.endpoint.v3.ClusterLoadAssignment
	(*v31.Cluster)(nil),                  // 3: envoy.config.cluster.v3.Cluster
	(*v32.RouteConfiguration)(nil),       // 4: envoy.config.route.v3.RouteConfiguration
	(*v32.ScopedRouteConfiguration)(nil), // 5: envoy.config.route.v3.ScopedRouteConfiguration
	(*v32.VirtualHost)(nil),              // 6: envoy.config.route.v3.VirtualHost
	(*v33.Listener)(nil),                 // 7: envoy.config.listener.v3.Listener
	(*v34.Secret)(nil),                   // 8: envoy.extensions.transport_sockets.tls.v3.Secret
	(*v35.TypedExtensionConfig)(nil),     // 9: envoy.config.core.v3.TypedExtensionConfig
	(*v36.Runtime)(nil),                  // 10: envoy.service.runtime.v3.Runtime
	(*v37.RouteConfiguration)(nil),       // 11: envoy.extensions.filters.network.thrift_proxy.v3.RouteConfiguration
}
var file_src_protoconfxds_v1_protoconfxds_proto_depIdxs = []int32{
	2,  // 0: protoconfenvoy.v1.XDSSnapshot.endpoints:type_name -> envoy.config.endpoint.v3.ClusterLoadAssignment
	3,  // 1: protoconfenvoy.v1.XDSSnapshot.clusters:type_name -> envoy.config.cluster.v3.Cluster
	4,  // 2: protoconfenvoy.v1.XDSSnapshot.routes:type_name -> envoy.config.route.v3.RouteConfiguration
	5,  // 3: protoconfenvoy.v1.XDSSnapshot.scoped_types:type_name -> envoy.config.route.v3.ScopedRouteConfiguration
	6,  // 4: protoconfenvoy.v1.XDSSnapshot.virtual_hosts:type_name -> envoy.config.route.v3.VirtualHost
	7,  // 5: protoconfenvoy.v1.XDSSnapshot.listeners:type_name -> envoy.config.listener.v3.Listener
	8,  // 6: protoconfenvoy.v1.XDSSnapshot.secrets:type_name -> envoy.extensions.transport_sockets.tls.v3.Secret
	9,  // 7: protoconfenvoy.v1.XDSSnapshot.extention_configs:type_name -> envoy.config.core.v3.TypedExtensionConfig
	10, // 8: protoconfenvoy.v1.XDSSnapshot.runtimes:type_name -> envoy.service.runtime.v3.Runtime
	11, // 9: protoconfenvoy.v1.XDSSnapshot.thrift_routes:type_name -> envoy.extensions.filters.network.thrift_proxy.v3.RouteConfiguration
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_src_protoconfxds_v1_protoconfxds_proto_init() }
func file_src_protoconfxds_v1_protoconfxds_proto_init() {
	if File_src_protoconfxds_v1_protoconfxds_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_protoconfxds_v1_protoconfxds_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProtoconfEnvoyConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_protoconfxds_v1_protoconfxds_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XDSSnapshot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_protoconfxds_v1_protoconfxds_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_protoconfxds_v1_protoconfxds_proto_goTypes,
		DependencyIndexes: file_src_protoconfxds_v1_protoconfxds_proto_depIdxs,
		MessageInfos:      file_src_protoconfxds_v1_protoconfxds_proto_msgTypes,
	}.Build()
	File_src_protoconfxds_v1_protoconfxds_proto = out.File
	file_src_protoconfxds_v1_protoconfxds_proto_rawDesc = nil
	file_src_protoconfxds_v1_protoconfxds_proto_goTypes = nil
	file_src_protoconfxds_v1_protoconfxds_proto_depIdxs = nil
}
