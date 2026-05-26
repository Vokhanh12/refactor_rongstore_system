package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	sqlc "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type DbAdapter struct {
	Q *sqlc.Queries
	P *pgxpool.Pool
}

func New(q *sqlc.Queries, p *pgxpool.Pool) *DbAdapter {
	return &DbAdapter{
		Q: q,
		P: p,
	}
}

func (db *DbAdapter) Wrap(err error) *aerrs.AppError {
	return TranslateDBError(err)
}
