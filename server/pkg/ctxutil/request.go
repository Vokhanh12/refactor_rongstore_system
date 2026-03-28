package ctxutil

import (
	"context"

	"github.com/google/uuid"
)

type requestCtxKeyType struct{}

var requestCtxKey = requestCtxKeyType{}

type RequestContext struct {
	RequestID string
	TraceID   string
	SpanID    string
	SessionID string
}

func NewRequestContext() RequestContext {
	return RequestContext{
		RequestID: uuid.NewString(),
		SessionID: uuid.NewString(),
	}
}

func WithRequest(ctx context.Context, r RequestContext) context.Context {
	return context.WithValue(ctx, requestCtxKey, r)
}

func Request(ctx context.Context) (RequestContext, bool) {
	v, ok := ctx.Value(requestCtxKey).(RequestContext)
	return v, ok
}
