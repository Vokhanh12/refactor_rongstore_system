package dispatcher

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type HandlerFunc func(
	ctx context.Context,
	payload any,
) (any, *aerrs.AppError)

type Dispatcher struct {
	handlers map[Operation]HandlerFunc
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: make(map[Operation]HandlerFunc),
	}
}

func (d *Dispatcher) Register(
	op Operation,
	handler HandlerFunc,
) *Dispatcher {

	if handler == nil {
		panic("dispatcher: nil handler")
	}

	if _, ok := d.handlers[op]; ok {
		panic("dispatcher: duplicate handler " + string(op))
	}

	d.handlers[op] = handler

	return d
}

func (d *Dispatcher) Dispatch(
	ctx context.Context,
	op Operation,
	payload any,
) (any, *aerrs.AppError) {

	handler, ok := d.handlers[op]

	if !ok {
		return nil, &core.HANDLER_NOT_FOUND
	}

	return handler(ctx, payload)
}
