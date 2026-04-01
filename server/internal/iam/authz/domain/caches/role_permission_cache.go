package caches

import (
	"context"
	"time"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type RolePermissionCache interface {
	GetPermissions(ctx context.Context, roleCode string, roleScopeId string) ([]string, *aerrs.AppError)

	SetPermissions(
		ctx context.Context,
		roleCode string,
		roleScopeId string,
		perms []string,
		ttl time.Duration,
	) *aerrs.AppError
}
