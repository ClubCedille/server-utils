package serverutils

import (
	"context"
	"net/http"
)

// HttpServer represents an instance
// of *http.Server, which gracefully shutdowns
// and has a status
type HttpServer struct {
	status Status
	server *http.Server
}

// NewHttpServer creates a new instance
// of *HttpServer
func NewHttpServer(server *http.Server) *HttpServer {
	return &HttpServer{server: server}
}

// Make sure struct implements interfaces.
var _ Server = &HttpServer{}
var _ serverOperations = &HttpServer{}

func (s *HttpServer) Run(ctx context.Context, req RunRequest) error {
	return startServer(ctx, s, req)
}

func (s *HttpServer) Status() Status {
	return s.status
}

func (s *HttpServer) gracefullyShutdown(ctx context.Context) error {
	// Update server status to Closed
	s.status = Stopped

	// Shutdown server
	return s.server.Shutdown(ctx)
}

func (s *HttpServer) serve(port int32) error {
	// Update server status to Started
	s.status = Running

	// Start server
	return s.server.ListenAndServe()
}
