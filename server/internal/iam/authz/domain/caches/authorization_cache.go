package caches

import (
	"context"
	"time"

	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type AuthorizationCache interface {
	GetResourceActionByRoleKey(ctx context.Context, RoleKey vo.RoleKey) ([]vo.ResourceAction, *aerrs.AppError)
	SetResourceActionByRoleKey(ctx context.Context, RoleKey vo.RoleKey, resourceActions []vo.ResourceAction, ttl time.Duration) *aerrs.AppError
}
