package serverutils

import "context"

type GrpcClient struct{}

func NewGrpcClient() *GrpcClient {
	return &GrpcClient{}
}

// Make sure struct implements interface.
var _ Machine = &GrpcClient{}

func (g *GrpcClient) Run(ctx context.Context, req MachineRequest) error {
	panic("implement me")
}

func (g *GrpcClient) Status() Status {
	panic("implement me")
}
