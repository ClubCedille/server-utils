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

type ConnectionRequest struct {
	Port       int32
	ShutdownMs int32
}

type Connection interface {
	Run(ctx context.Context, req ConnectionRequest) error

	Status() Status
}
