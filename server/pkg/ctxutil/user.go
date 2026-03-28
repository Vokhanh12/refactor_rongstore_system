package ctxutil

import "context"

type UserContext struct {
	UserID string
	Roles  []string
}

type userCtxKeyType struct{}

var userCtxKey = userCtxKeyType{}

func WithUser(ctx context.Context, user UserContext) context.Context {
	return context.WithValue(ctx, userCtxKey, user)
}

func User(ctx context.Context) (UserContext, bool) {
	v, ok := ctx.Value(userCtxKey).(UserContext)
	return v, ok
}
