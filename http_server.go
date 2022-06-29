package serverutils

import (
	"context"
	"net/http"
)

type HttpServer struct {
	status Status
	server *http.Server
}

func NewHttpServer(server *http.Server) *HttpServer {
	return &HttpServer{server: server}
}

// Make sure struct implements interface.
var _ Connection = &HttpServer{}

func (s *HttpServer) Run(ctx context.Context, req ConnectionRequest) error {
	return startServer(ctx, s, req)
}

func (s *HttpServer) Status() Status {
	panic("implement me")
}

func (s *HttpServer) gracefulShutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *HttpServer) serve(port int32) error {
	return s.server.ListenAndServe()
}
