package serverutils

import "context"

type GrpcServer struct{}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

// Make sure struct implements interface.
var _ Machine = &GrpcServer{}

func (g *GrpcServer) Run(ctx context.Context, req MachineRequest) error {
	panic("implement me")
}

func (g *GrpcServer) Status() Status {
	panic("implement me")
}
