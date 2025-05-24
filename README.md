# Go gRPC Example Project

This project demonstrates how to build a gRPC server and client in Go, using Protocol Buffers for message serialization. It covers unary, server streaming, client streaming, and bidirectional streaming RPCs.

## Project Structure

```
Go_grpc/
├── client/
│   └── main.go         # gRPC client implementation
├── proto/
│   ├── greet.proto     # Protobuf service and message definitions
│   └── greet.pb.go     # Generated Go code from greet.proto
├── server/
│   └── main.go         # gRPC server implementation
└── go.mod              # Go module definition
```

## High-Level Concepts

- **gRPC:** A high-performance, open-source RPC framework that uses HTTP/2 and Protocol Buffers.
- **Protocol Buffers (protobuf):** A language-neutral, platform-neutral extensible mechanism for serializing structured data.
- **Unary RPC:** Client sends a single request and gets a single response.
- **Server Streaming RPC:** Client sends a request and receives a stream of responses.
- **Client Streaming RPC:** Client sends a stream of requests and receives a single response.
- **Bidirectional Streaming RPC:** Both client and server send a stream of messages to each other.

## How It Works

- The `proto/greet.proto` file defines the gRPC service (`GreetService`) and messages.
- The server (`server/main.go`) implements the service methods.
- The client (`client/main.go`) demonstrates calling each type of RPC.
- Protobuf files are compiled to Go code using `protoc`.

## Getting Started

### Prerequisites

- Go 1.18+
- `protoc` Protocol Buffers compiler
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

### Install Dependencies

```sh
go mod tidy
```

### Generate Go Code from Protobuf

```sh
protoc --go_out=. --go-grpc_out=. proto/greet.proto
```

### Run the Server

```sh
go run server/main.go
```

### Run the Client

```sh
go run client/main.go
```

## Customization

- Edit `proto/greet.proto` to change service or message definitions.
- Re-run the `protoc` command after making changes to regenerate Go code.

## References

- [gRPC in Go](https://grpc.io/docs/languages/go/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
