package cache

import (
	"context"
	"time"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RolePermissionCache interface {
	GetPermissions(ctx context.Context, roleCode string) ([]string, bool, *aerrs.AppError)

	SetPermissions(
		ctx context.Context,
		roleCode string,
		perms []string,
		ttl time.Duration,
	) *aerrs.AppError
}
