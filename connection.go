package serverutils

import (
	"context"
)

type Status int

const (
	StatusStarted Status = iota
	StatusClosed
	StatusFailed
)

type ConnectionRequest struct {
	Port              int32
	ShutdownTimeoutMs int32
}

type Connection interface {
	Run(ctx context.Context, req ConnectionRequest) error
	Status() Status
}
