package caches

import (
	"context"
	"time"

	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RolePermissionCache interface {
	GetPermissions(ctx context.Context, roleRef vo.RoleRef) ([]vo.ResourceAction, *aerrs.AppError)
	SetPermissions(ctx context.Context, roleRef vo.RoleRef, resourceActions []vo.ResourceAction, ttl time.Duration) *aerrs.AppError
}
