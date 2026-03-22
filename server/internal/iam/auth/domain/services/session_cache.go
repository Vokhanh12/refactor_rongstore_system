package services

import (
	"context"

	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/entities"
)

type SessionCache interface {
	//	CheckHealth(r *RedisSessionStore) *errors.BusinessError
	StoreSession(ctx context.Context, e *en.SessionEntry) error
	GetSession(ctx context.Context, sessionID string) (*en.SessionEntry, error)
	DeleteSession(ctx context.Context, sessionID string) error
	CheckAndRecordNonceAtomic(ctx context.Context, sessionID, nonceB64 string, windowSeconds int) (bool, error)
}
