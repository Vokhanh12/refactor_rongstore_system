package errors

import (
	"context"
	"database/sql"
	"errors"

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
		}
	}

	// ---------- Timeout ----------
	if errors.Is(err, context.DeadlineExceeded) {
		return aerrs.New(
			domain.DB_TIMEOUT,
			aerrs.WithCauseDetail(err),
		)
	}

	// ---------- Default ----------
	return aerrs.New(
		catalogs.PersistCatalog,
		aerrs.WithCauseDetail(err),
	)
}
