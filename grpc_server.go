package serverutils

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	status Status
	server *grpc.Server
}

func NewGrpcServer(server *grpc.Server) *GrpcServer {
	return &GrpcServer{server: server}
}

// Make sure struct implements interface.
var _ Connection = &GrpcServer{}

func (g *GrpcServer) Run(ctx context.Context, req ConnectionRequest) error {
	return startServer(ctx, g, req)
}

func (g *GrpcServer) Status() Status {
	panic("implement me")
}

func (g *GrpcServer) serve(port int32) error {
	// Start a new connection on given port
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %s", port, err)
	}

	// Serve gRPC server
	err = g.server.Serve(conn)
	if err != nil {
		return fmt.Errorf("failed to serve gRPC connection: %s", err)
	}

	// No error occured, exit
	return nil
}

func (g *GrpcServer) gracefulShutdown(ctx context.Context) error {
	doneCh := make(chan bool, 1)
	defer close(doneCh)
	go func() {
		g.server.GracefulStop()
		doneCh <- true
	}()
	select {
	case <-ctx.Done():
	case <-doneCh:
	}
	return nil
}
