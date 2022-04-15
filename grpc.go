package serverutils

import "context"

type GrpcServer struct{}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

// Make sure struct implements interface.
var _ Server = &GrpcServer{}

func (g *GrpcServer) Run(ctx context.Context, opts Request) error {
	panic("implement me")
}

func (g *GrpcServer) Status() Status {
	panic("implement me")
}
