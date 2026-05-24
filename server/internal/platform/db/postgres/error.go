package postgres

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	plerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

	"github.com/jackc/pgx/v5/pgconn"
)

func TranslateDBError(err error) *aerrs.AppError {
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return aerrs.New(plerrs.DB_NOT_FOUND)
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {

		switch pgErr.Code {

		case "23505":
			return aerrs.New(plerrs.DB_CONFLICT)

		case "23503", "23514":
			return aerrs.New(plerrs.DB_INVALID_REFERENCE)

		case "57014":
			return aerrs.New(plerrs.DB_TIMEOUT)
		}
	}

	if errors.Is(err, context.DeadlineExceeded) {
		return aerrs.New(plerrs.DB_TIMEOUT)
	}

	if isConnectionError(err) {
		return aerrs.New(plerrs.POSTGRES_UNAVAILABLE)
	}

	return aerrs.New(plerrs.DB_QUERY_FAILED)
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
