package rediscache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/cache"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/caches"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	errs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/errors"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

var _ caches.AuthorizationCache = (*RedisAuthorizationCache)(nil)

type RedisAuthorizationCache struct {
	client *redis.Client
	codec  core.Codec
}

// GetResourceActionByRoleKey implements [caches.AuthorizationCache].
func (r *RedisAuthorizationCache) GetResourceActionByRoleKey(ctx context.Context, RoleKey vo.RoleKey) ([]vo.ResourceAction, *aerrs.AppError) {
	panic("unimplemented")
}

// GetResourceAction implements [caches.AuthorizationCache].
func (r *RedisAuthorizationCache) GetResourceAction(ctx context.Context, RoleKey valueobjects.RoleKey) ([]valueobjects.ResourceAction, *apperrors.AppError) {

	val, err := r.client.Get(ctx, r.cacheKey(RoleKey.String())).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, aerrs.New(errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err))
	}

	var resourceActions []vo.ResourceAction
	if err := json.Unmarshal([]byte(val), &resourceActions); err != nil {
		return nil, aerrs.New(errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err))
	}

	return resourceActions, nil
}

// SetResourceActionByRoleKey implements [caches.AuthorizationCache].
func (r *RedisAuthorizationCache) SetResourceActionByRoleKey(ctx context.Context, RoleKey valueobjects.RoleKey, resourceActions []valueobjects.ResourceAction, ttl time.Duration) *apperrors.AppError {
	data, err := json.Marshal(resourceActions)
	if err != nil {
		return aerrs.New(errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err))
	}

	if err := r.client.Set(ctx, r.cacheKey(RoleKey.String()), data, ttl).Err(); err != nil {
		return aerrs.New(errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err))
	}

	return nil
}

func (r *RedisAuthorizationCache) cacheKey(value string) string {
	return "authz:authorization:" + value
}
