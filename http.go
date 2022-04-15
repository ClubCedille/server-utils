package serverutils

import "context"

type HttpServer struct {
	status Status
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

// Make sure struct implements interface.
var _ Server = &HttpServer{}

func (g *HttpServer) Run(ctx context.Context, opts Request) error {
	panic("implement me")
}

func (g *HttpServer) Status() Status {
	panic("implement me")
}
