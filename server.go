package serverutils

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type server interface {
	gracefulShutdown(ctx context.Context) error
	serve(port int32) error
}

// Make sure structs implement
// the Server interface
var _ server = &GrpcServer{}
var _ server = &HttpServer{}

func startServer(ctx context.Context, s server, req ConnectionRequest) error {
	ctx, cancel := context.WithTimeout(
		ctx,
		time.Millisecond*time.Duration(req.ShutdownTimeoutMs),
	)
	defer cancel()

	// Catch some signals and handle them gracefully
	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		sig := <-sigCh
		log.Printf("got signal %v, attempting graceful shutdown\n", sig)
		cancel()

		// Pass returned error from shutdown
		// to error channel
		errCh <- s.gracefulShutdown(ctx)

		// grpc.Stop() // leads to error while receiving stream response: rpc error: code = Unavailable desc = transport is closing
		wg.Done()
	}()

	// Close channel
	close(errCh) // TODO: this might be wrong, make sure errCh it properly closed

	// Catch error and return it
	select {
	case <-ctx.Done():
	case err := <-errCh:
		return fmt.Errorf("failed to shutdown server: %s", err)
	}

	// Start the gRPC server
	err := s.serve(req.Port)
	if err != nil {
		return fmt.Errorf("failed to serve connection: %s", err)
	}
	wg.Wait()

	return nil
}
