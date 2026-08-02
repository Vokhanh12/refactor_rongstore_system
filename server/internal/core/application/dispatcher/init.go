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
	handlers map[Action]HandlerFunc
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers: make(map[Action]HandlerFunc),
	}
}

func (d *Dispatcher) Register(
	action Action,
	handler HandlerFunc,
) *Dispatcher {

	if handler == nil {
		panic("dispatcher: nil handler")
	}

	if _, ok := d.handlers[action]; ok {
		panic(
			"dispatcher: duplicate handler " +
				string(action),
		)
	}

	d.handlers[action] = handler

	return d
}

func (d *Dispatcher) Dispatch(
	ctx context.Context,
	action Action,
	payload any,
) (any, *aerrs.AppError) {

	handler, ok := d.handlers[action]

	if !ok {
		return nil, &core.HANDLER_NOT_FOUND
	}

	return handler(
		ctx,
		payload,
	)
}
