package rediscache

import (
	"errors"

	plerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var (
	ErrSerialize        = errors.New("cache serialize error")
	ErrDeserialize      = errors.New("cache deserialize error")
	ErrRedisUnavailable = errors.New("redis unavailable")
)

func TranslateCacheError(
	err error,
) *aerrs.AppError {

	switch {

	case errors.Is(err, ErrSerialize):
		return aerrs.New(
			plerrs.CACHE_SERIALIZATION_FAILED,
			aerrs.WithCauseDetail(err),
		)

	case errors.Is(err, ErrDeserialize):
		return aerrs.New(
			plerrs.CACHE_DESERIALIZATION_FAILED,
			aerrs.WithCauseDetail(err),
		)

	case errors.Is(err, ErrRedisUnavailable):
		return aerrs.New(
			plerrs.REDIS_UNAVAILABLE,
			aerrs.WithCauseDetail(err),
		)
	}

	return aerrs.New(
		aerrs.INTERNAL_FALLBACK,
		aerrs.WithCauseDetail(err),
	)
}
