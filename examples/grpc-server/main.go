package main

import (
	"context"
	"log"

	serverutils "github.com/clubcedille/server-utils"
	"google.golang.org/grpc"
)

func main() {
	// Define your own gRPC server
	server := grpc.NewServer()
	// register server with protobuf...

	// Create and run new gRPC server
	grpcServer := serverutils.NewGrpcServer(server)

	// Run newly created gRPC server
	req := serverutils.RunRequest{
		Port:              3000,
		ShutdownTimeoutMs: 100000,
	}
	if err := grpcServer.Run(context.Background(), req); err != nil {
		log.Fatalf("fatal error when starting gRPC server: %s\n", err)
	}
}
