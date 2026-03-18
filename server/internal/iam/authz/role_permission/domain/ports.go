package domain

import (
	"context"
	pkg "server/pkg/apperrors"
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

type RolePermissionRepository interface {
	GetPermissions(ctx context.Context, roleCode string) ([]string, bool, *pkg.AppError)

	SetPermissions(
		ctx context.Context,
		roleCode string,
		perms []string,
		ttl time.Duration,
	) *pkg.AppError
}
