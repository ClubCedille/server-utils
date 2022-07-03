package serverutils

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// GrpcServer represents an instance
// of *grpc.Server, which gracefully shutdowns
// and has a status
type GrpcServer struct {
	status Status
	server *grpc.Server
}

// NewGrpcServer creates a new instance
// of *GrpcServer
func NewGrpcServer(server *grpc.Server) *GrpcServer {
	return &GrpcServer{server: server}
}

// Make sure struct implements interface.
var _ Server = &GrpcServer{}
var _ serverOperations = &GrpcServer{}

func (g *GrpcServer) Run(ctx context.Context, req RunRequest) error {
	return startServer(ctx, g, req)
}

func (g *GrpcServer) Status() Status {
	return g.status
}

func (g *GrpcServer) serve(port int32) error {
	// Start a new connection on given port
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %s", port, err)
	}

	// Update server status to Started
	g.status = Running

	// Serve gRPC server
	err = g.server.Serve(conn)
	if err != nil {
		return fmt.Errorf("failed to serve gRPC connection: %s", err)
	}

	// No error occured, exit
	return nil
}

func (g *GrpcServer) gracefullyShutdown(ctx context.Context) error {
	doneCh := make(chan bool, 1)
	go func() {
		g.server.GracefulStop()
		doneCh <- true

		// Update server status to Closed
		g.status = Stopped
	}()

	select {
	case <-ctx.Done():
	case <-doneCh:
	}

	return nil
}
