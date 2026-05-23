// ============================================================
// AUTHZ REDIS ADAPTER
// internal/iam/authz/infrastructure/caches/rediscache/authorization_cache.go
// ============================================================

package rediscache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/caches"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/cache"
	cache "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/cache"
	cache "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/cache/errors"
	errs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/errors"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ caches.AuthorizationCache = (*RedisAuthorizationCache)(nil)

const authorizationPrefix = "authz:role-permissions"

type RedisAuthorizationCache struct {
	cache      *cache.RedisCache
	keyBuilder *core.KeyBuilder
}

func NewRedisAuthorizationCache(client *redis.Client) *RedisAuthorizationCache {

	return &RedisAuthorizationCache{
		cache: cache.NewRedisCache(
			client,
			core.NewJSONCodec(),
		),

		keyBuilder: core.NewKeyBuilder(
			authorizationPrefix,
		),
	}
}

func (r *RedisAuthorizationCache) GetResourceActionByRoleKey(
	ctx context.Context,
	roleKey vo.RoleKey,
) ([]vo.ResourceAction, *aerrs.AppError) {

	key := r.buildRoleKey(roleKey)

	var resourceActions []vo.ResourceAction

	err := r.cache.Get(
		ctx,
		key,
		&resourceActions,
	)

	if err != nil {

		if err == redis.Nil {
			return nil, nil
		}

		return nil, aerrs.New(
			errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err),
		)
	}

	return resourceActions, nil
}

func (r *RedisAuthorizationCache) SetResourceActionByRoleKey(
	ctx context.Context,
	roleKey vo.RoleKey,
	resourceActions []vo.ResourceAction,
	ttl time.Duration,
) *aerrs.AppError {

	key := r.buildRoleKey(roleKey)

	err := r.cache.Set(
		ctx,
		key,
		resourceActions,
		ttl,
	)

	if err != nil {
		return aerrs.New(
			errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err),
		)
	}

	return nil
}

func (r *RedisAuthorizationCache) DeleteByRoleKey(
	ctx context.Context,
	roleKey vo.RoleKey,
) *aerrs.AppError {

	key := r.buildRoleKey(roleKey)

	err := r.cache.Delete(
		ctx,
		key,
	)

	if err != nil {
		return aerrs.New(
			errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err),
		)
	}

	return nil
}

func (r *RedisAuthorizationCache) buildRoleKey(
	roleKey vo.RoleKey,
) string {

	return r.keyBuilder.Build(
		roleKey.String(),
	)
}
