package postgres

import (
	sqlc "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type DB struct {
	Q *sqlc.Queries
}

func New(q *sqlc.Queries) *DB {
	return &DB{
		Q: q,
	}
}

func (db *DB) Wrap(err error) *aerrs.AppError {
	return TranslateDBError(err)
}
