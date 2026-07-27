package application

import (
	"context"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type HandlerFunc func(ctx context.Context, payload any) (any, *aerrs.AppError)

type Dispatcher struct {
	handlers map[string]HandlerFunc
}

func NewDispatcher() *Dispatcher {

	return &Dispatcher{
		handlers: make(map[string]HandlerFunc),
	}

}

func (d *Dispatcher) Register(operation string, handler HandlerFunc) {

	d.handlers[operation] = handler

}

func (d *Dispatcher) Dispatch(
	ctx context.Context,

	operation string,

	payload any,

) (any, *aerrs.AppError) {

	handler, ok := d.handlers[operation]

	if !ok {
		return nil, ErrHandlerNotFound
	}

	return handler(ctx, payload)

}
