package ctxutil

import (
	"context"

	"github.com/google/uuid"
)

type IdentityContext struct {
	UserID       string
	RoleScopes   []RoleScope
	AuthzVersion int
}

type RoleScope struct {
	RoleID    uuid.UUID
	ScopeID   *uuid.UUID
	ScopeType string
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
