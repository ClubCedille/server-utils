# server-utils

Utils library for Go HTTP & gRPC servers.

## Install this library

To use this library in your own project, simply run this command:
```bash
$ go get github.com/clubcedille/server-utils@latest
...
```

## gRPC Server

The `server-utils` library offers a gRPC implementation of graceful shutdown and real-time status fetching.
See [this example](./examples/grpc-server/main.go) to learn how to use its functions.

## HTTP Server

The `server-utils` library offers an HTTP implementation of graceful shutdown and real-time status fetching.
See [this example](./examples/http-server/main.go) to learn how to use its functions.