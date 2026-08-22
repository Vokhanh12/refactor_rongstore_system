package test

import (
	"context"

	sec "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
)

type IdentityContext struct {
	UserID       string
	RoleScopes   []sec.RoleScope
	AuthzVersion int
}

type identityCtxKeyType struct{}

var identityCtxKey = identityCtxKeyType{}

func WithIdentity(ctx context.Context, identity IdentityContext) context.Context {
	return context.WithValue(ctx, identityCtxKey, identity)
}

func Identity(ctx context.Context) (IdentityContext, bool) {
	v, ok := ctx.Value(identityCtxKey).(IdentityContext)
	return v, ok
}
