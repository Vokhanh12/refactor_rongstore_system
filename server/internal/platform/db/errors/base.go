package errors

import (
	"context"
	"database/sql"
	"errors"

	"server/internal/iam/domain"
	pkgerrors "server/pkg/errors"

	"github.com/jackc/pgx/v5/pgconn"
)

type DBError struct {
	InvalidCatalog  pkgerrors.AppError
	NotfoundCatalog pkgerrors.AppError
	ConflictCatalog pkgerrors.AppError
	PersistCatalog  pkgerrors.AppError
}

func TranslateDBError(err error, catalogs DBError) *pkgerrors.AppError {
	if err == nil {
		return nil
	}

	// ---------- Not found ----------
	if errors.Is(err, sql.ErrNoRows) {
		return pkgerrors.New(
			catalogs.NotfoundCatalog,
			pkgerrors.WithCauseDetail(err),
		)
	}

	// ---------- Postgres specific ----------
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {

		switch pgErr.Code {

		case "23505": // unique_violation
			return pkgerrors.New(
				catalogs.ConflictCatalog,
				pkgerrors.WithCauseDetail(err),
			)

		case "23503", "23514": // FK, CHECK
			return pkgerrors.New(
				catalogs.InvalidCatalog,
				pkgerrors.WithCauseDetail(err),
			)
		}
	}

	// ---------- Timeout ----------
	if errors.Is(err, context.DeadlineExceeded) {
		return pkgerrors.New(
			domain.DB_TIMEOUT,
			pkgerrors.WithCauseDetail(err),
		)
	}

	// ---------- Default ----------
	return pkgerrors.New(
		catalogs.PersistCatalog,
		pkgerrors.WithCauseDetail(err),
	)
}
