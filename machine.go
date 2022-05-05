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

type MachineRequest struct {
	Port       int32
	ShutdownMs int32
}

type Machine interface {
	Run(ctx context.Context, req MachineRequest) error

	Status() Status
}
