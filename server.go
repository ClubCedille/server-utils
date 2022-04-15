package serverutils

import (
	"context"
)

type Status int

const (
	StatusClosed Status = iota
	StatusStarted
	StatusFailed
	StatusWhatever // TODO: Remove this
)

type Request struct {
	Port       int32
	ShutdownMs int32
}

type Server interface {
	Run(ctx context.Context, opts Request) error

	Status() Status
}
