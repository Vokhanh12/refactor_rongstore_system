package redisprovider

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/config"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/errors"
)

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
