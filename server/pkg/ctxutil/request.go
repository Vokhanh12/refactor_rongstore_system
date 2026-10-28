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
	ClientIP  string // 👈 Bổ sung IP
	UserAgent string // 👈 Bổ sung User-Agent
	TenantID  string // 👈 Bổ sung TenantID
}

func NewRequestContext() RequestContext {
	return RequestContext{
		RequestID: uuid.NewString(),
		// SessionID thường sẽ được gán sau khi Authenticated,
		// hoặc giữ nguyên nếu bạn sinh sẵn session từ đầu.
	}
}

func WithRequest(ctx context.Context, r RequestContext) context.Context {
	return context.WithValue(ctx, requestCtxKey, r)
}

func Request(ctx context.Context) (RequestContext, bool) {
	v, ok := ctx.Value(requestCtxKey).(RequestContext)
	return v, ok
}

func MustRequest(ctx context.Context) RequestContext {
	v, ok := ctx.Value(requestCtxKey).(RequestContext)
	if !ok {
		panic("request not found in context")
	}
	return v
}
