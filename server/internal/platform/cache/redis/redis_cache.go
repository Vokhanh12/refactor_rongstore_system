package rediscache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/cache"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/config"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/errors"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RedisCache struct {
	client *redis.Client
	codec  core.Codec
}

func NewRedisClient(
	cfg *config.Config,
) (*redis.Client, *apperrors.AppError) {

	addr := fmt.Sprintf(
		"%s:%d",
		cfg.RedisHost,
		cfg.RedisPort,
	)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,

		PoolSize:     cfg.RedisPoolSize,
		MinIdleConns: cfg.RedisMinIdleConns,

		DialTimeout:  3 * time.Second,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,

		MaxRetries: 3,
	})

	ctx, cancel := context.WithTimeout(
		context.Background(),
		2*time.Second,
	)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {

		_ = rdb.Close()

		return nil, apperrors.New(
			errors.REDIS_UNAVAILABLE,
			apperrors.WithCauseDetail(err),
		)
	}

	return rdb, nil
}

func NewRedisCache(
	client *redis.Client,
	codec core.Codec,
) *RedisCache {

	return &RedisCache{
		client: client,
		codec:  codec,
	}
}

func (r *RedisCache) Get(
	ctx context.Context,
	key string,
	dest any,
) error {

	data, err := r.client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}

	return r.codec.Unmarshal(data, dest)
}

func (r *RedisCache) Set(
	ctx context.Context,
	key string,
	value any,
	ttl time.Duration,
) error {

	data, err := r.codec.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, ttl).Err()
}

func (r *RedisCache) Delete(
	ctx context.Context,
	key string,
) error {

	return r.client.Del(ctx, key).Err()
}
