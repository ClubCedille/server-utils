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

type Status int

const (
	// Running is the status
	// for a running server
	Running Status = iota

	// Stopped is the status
	// for a stopped server
	Stopped
)

// RunRequest is a set of parameters
// passed to the 'Run' function of each
// server
type RunRequest struct {
	// Port is the port number on which
	// the server will listen to
	Port int32

	// ShutdownTimeoutMs is the amount
	// of milliseconds to wait for when
	// shutting down the server
	ShutdownTimeoutMs int32
}

// Server represents any kind of server
// that is able to run, shutdown and
// fetch a status from.
type Server interface {
	// Run runs the server on a given port
	// and gracefully shutdowns if necessary
	Run(ctx context.Context, req RunRequest) error

	// Status returns the current status
	// of the server
	Status() Status
}

// serverOperations is an internal interface meant
// to serve as a helper for server functions.
type serverOperations interface {
	gracefullyShutdown(ctx context.Context) error
	serve(port int32) error
}

func startServer(ctx context.Context, s serverOperations, req RunRequest) error {
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
		errCh <- s.gracefullyShutdown(ctx)

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
