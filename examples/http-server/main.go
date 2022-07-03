package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	serverutils "github.com/clubcedille/server-utils"
)

func main() {
	// Define your own gRPC server
	port := 3000
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	// Create and run new gRPC server
	httpServer := serverutils.NewHttpServer(server)

	// Run newly created gRPC server
	req := serverutils.RunRequest{
		Port:              int32(port),
		ShutdownTimeoutMs: 100000,
	}
	if err := httpServer.Run(context.Background(), req); err != nil {
		log.Fatalf("fatal error when starting http server: %s\n", err)
	}
}
