package rediscache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/caches"
	errs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

var _ caches.RolePermissionCache = (*RedisRolePermissionCache)(nil)

type RedisRolePermissionCache struct {
	client *redis.Client
}

func (r *RedisRolePermissionCache) GetPermissions(ctx context.Context, roleCode string) ([]string, *aerrs.AppError) {

	val, err := r.client.Get(ctx, r.key(roleCode)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, aerrs.New(errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err))
	}

	var perms []string
	if err := json.Unmarshal([]byte(val), &perms); err != nil {
		return nil, aerrs.New(errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err))
	}

	return perms, nil
}

func (r *RedisRolePermissionCache) SetPermissions(ctx context.Context, roleCode string, perms []string, ttl time.Duration) *aerrs.AppError {

	data, err := json.Marshal(perms)
	if err != nil {
		return aerrs.New(errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err))
	}

	if err := r.client.Set(ctx, r.key(roleCode), data, ttl).Err(); err != nil {
		return aerrs.New(errs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err))
	}

	return nil
}

func (r *RedisRolePermissionCache) key(roleCode string) string {
	return "authz:role_permission:" + roleCode
}
