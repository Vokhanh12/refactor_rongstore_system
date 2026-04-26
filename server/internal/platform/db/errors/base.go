package errors

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	plerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

	"github.com/jackc/pgx/v5/pgconn"
)

type DBError struct {
	// ---------- Business-like ----------
	Conflict aerrs.AppError // 23505 (unique)
	Invalid  aerrs.AppError // FK, CHECK
	NotFound aerrs.AppError // sql.ErrNoRows

	// ---------- Infra ----------
	//Timeout     aerrs.AppError // context deadline
	//Unavailable aerrs.AppError // connection issue
	Internal aerrs.AppError // fallback (unknown)
}

func TranslateDBError(err error, e DBError) *aerrs.AppError {
	if err == nil {
		return nil
	}

	// ---------- Not Found ----------
	if errors.Is(err, sql.ErrNoRows) {
		return aerrs.New(e.NotFound, aerrs.WithCauseDetail(err))
	}

	// ---------- Postgres ----------
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {

		switch pgErr.Code {

		case "23505": // unique_violation
			return aerrs.New(e.Conflict, aerrs.WithCauseDetail(err))

		case "23503", "23514": // FK, CHECK
			return aerrs.New(e.Invalid, aerrs.WithCauseDetail(err))

		case "57014": // query canceled / timeout
			return aerrs.New(plerrs.DB_TIMEOUT, aerrs.WithCauseDetail(err))
		}
	}

	// ---------- Context timeout ----------
	if errors.Is(err, context.DeadlineExceeded) {
		return aerrs.New(plerrs.DB_TIMEOUT, aerrs.WithCauseDetail(err))
	}

	// ---------- Connection ----------
	if isConnectionError(err) {
		return aerrs.New(plerrs.POSTGRES_UNAVAILABLE, aerrs.WithCauseDetail(err))
	}

	// ---------- Fallback ----------
	return aerrs.New(e.Internal, aerrs.WithCauseDetail(err))
}

func isConnectionError(err error) bool {
	if errors.Is(err, context.Canceled) {
		return true
	}

	msg := err.Error()

	return strings.Contains(msg, "connection refused") ||
		strings.Contains(msg, "connection reset") ||
		strings.Contains(msg, "broken pipe") ||
		strings.Contains(msg, "EOF")
}
