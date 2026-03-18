package ports

import (
	"context"
	"time"
)

type RolePermissionCache interface {
	GetPermissions(ctx context.Context, roleCode string) ([]string, bool, *pkg.AppError)

	SetPermissions(
		ctx context.Context,
		roleCode string,
		perms []string,
		ttl time.Duration,
	) *pkg.AppError
}
