package errors

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"

	"github.com/jackc/pgx/v5/pgconn"
)

type DBError struct {
	InvalidCatalog  aerrs.AppError
	NotfoundCatalog aerrs.AppError
	ConflictCatalog aerrs.AppError
	PersistCatalog  aerrs.AppError
}

func TranslateDBError(err error, catalogs DBError) *aerrs.AppError {
	if err == nil {
		return nil
	}

	// ---------- Not found ----------
	if errors.Is(err, sql.ErrNoRows) {
		return aerrs.New(
			catalogs.NotfoundCatalog,
			aerrs.WithCauseDetail(err),
		)
	}

	// ---------- Postgres specific ----------
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {

		switch pgErr.Code {

		case "23505": // unique_violation
			return aerrs.New(
				catalogs.ConflictCatalog,
				aerrs.WithCauseDetail(err),
			)

		case "23503", "23514": // FK, CHECK
			return aerrs.New(
				catalogs.InvalidCatalog,
				aerrs.WithCauseDetail(err),
			)

		case "57014": // query_canceled (statement_timeout)
			return aerrs.New(
				domain.DB_QUERY_TIMEOUT,
				aerrs.WithCauseDetail(err),
			)
		}
	}

	// ---------- Timeout (context) ----------
	if errors.Is(err, context.DeadlineExceeded) {
		// ⚠️ tricky: không biết chắc là query hay connect
		// => assume là DB unavailable (an toàn hơn)
		return aerrs.New(
			domain.DB_TIMEOUT,
			aerrs.WithCauseDetail(err),
		)
	}

	// ---------- Connection issues ----------
	if isConnectionError(err) {
		return aerrs.New(
			domain.POSTGRES_UNAVAILABLE,
			aerrs.WithCauseDetail(err),
		)
	}

	// ---------- Default ----------
	return aerrs.New(
		catalogs.PersistCatalog,
		aerrs.WithCauseDetail(err),
	)
}

func isConnectionError(err error) bool {
	return errors.Is(err, context.Canceled) ||
		strings.Contains(err.Error(), "connection refused") ||
		strings.Contains(err.Error(), "connection reset") ||
		strings.Contains(err.Error(), "broken pipe")
}
