# Protoconf XDS Server

Protoconf XDS is an implementation of the Envoy's External Data Service (XDS) API that uses Protoconf as a configuration backend. It provides an easy way to manage and update Envoy's configuration dynamically.

## Getting Started

### Prerequisites

- Go 1.16 or later
- Protoconf v0.1.7 or later

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/protoconf/protoconf-xds.git
   ```

2. Build the binary:

   ```bash
   cd protoconf-xds
   go build
   ```

### Usage

Start the Protoconf XDS server by running:

```bash
./protoconf-xds
```

By default, the server listens on port 18000. You can change the port by specifying the `-port` flag:

```bash
./protoconf-xds -port 8080
```

### Configuration

Protoconf XDS reads its configuration from the environment variables or command line:

```
-debug
      env key: PROTOCONF_XDS_DEBUG
      type: bool (default false)
-nodeId value (repeatable)
      env key: PROTOCONF_XDS_NODE_ID (comma separate for multiple node ids)
      type: string (default [])
-port value
      env key: PROTOCONF_XDS_PORT
      type: uint32 (default 18000)
-prefix value
      env key: PROTOCONF_XDS_PREFIX
      type: string (default example)
-protoconfAgentAddr value
      env key: PROTOCONF_XDS_PROTOCONF_AGENT_ADDR
      type: string (default localhost:4300)
```

### License

Protoconf XDS is licensed under the MIT License. See `LICENSE` file for more information.

### Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

### Acknowledgments

This project was inspired by the [Envoy xDS server example](https://github.com/envoyproxy/go-control-plane/tree/main/examples) and uses [Protoconf gRPC Client](https://pkg.go.dev/github.com/protoconf/protoconf@v0.1.6/agent/api/proto/v1) to communicate with the Protoconf API server.
