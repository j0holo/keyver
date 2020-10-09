# Keyver

A persistent key-value store for a hobby project. It uses gRPC for communication with the client.

## Build the protocol.pb.go file

[Install protoc](https://github.com/protocolbuffers/protobuf/releases)

[Install protoc-gen-go](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers)

```bash
# Make sure protoc and protoc-gen-go are installed and in your path
protoc --go_opt=paths=source_relative --go_out=plugins=grpc:. rpc/protocol.proto
```

## Development

From the project root directory you can run unit tests and benchmarks.

Run tests with:

```bash
go test ./...
```

Run benchmarks with:

```bash
go test -bench=. ./...
```