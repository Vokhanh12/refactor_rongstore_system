package rediscache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/caches"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	errs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ caches.RolePermissionCache = (*RedisRolePermissionCache)(nil)

type RedisRolePermissionCache struct {
	client *redis.Client
}

func (r *RedisRolePermissionCache) GetPermissions(ctx context.Context, roleRef vo.RoleRef) ([]vo.ResourceAction, *aerrs.AppError) {

	val, err := r.client.Get(ctx, r.cacheKey(roleRef.String())).Result()
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

func (r *RedisRolePermissionCache) SetPermissions(ctx context.Context, roleRef vo.RoleRef, resourceActions []vo.ResourceAction, ttl time.Duration) *aerrs.AppError {

	data, err := json.Marshal(resourceActions)
	if err != nil {
		return aerrs.New(errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err))
	}

	if err := r.client.Set(ctx, r.cacheKey(roleRef.String()), data, ttl).Err(); err != nil {
		return aerrs.New(errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err))
	}

	return nil
}

func (r *RedisRolePermissionCache) cacheKey(value string) string {
	return "authz:role_permission:" + value
}
