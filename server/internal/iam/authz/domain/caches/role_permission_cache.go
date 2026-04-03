package caches

import (
	"context"
	"time"

	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type RolePermissionCache interface {
	GetPermissions(ctx context.Context, roles vo.RoleRef) ([]vo.ResourceAction, *aerrs.AppError)

	SetPermissions(
		ctx context.Context,
		roleKey vo.RoleRef,
		permKeys []vo.ResourceAction,
		ttl time.Duration,
	) *aerrs.AppError
}
