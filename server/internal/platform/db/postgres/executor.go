package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	sq "github.com/Masterminds/squirrel"
	plerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/errors"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type Scanner[T any] func(pgx.Rows) (T, error)

func QueryMany[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	sql string,
	args []any,
	scan Scanner[T],
) ([]T, *apperrors.AppError) {

	rows, err := db.Query(ctx, sql, args...)

	if err != nil {
		return nil, translateDBError(err)
	}

	defer rows.Close()

	results := make([]T, 0)

	for rows.Next() {

		item, err := scan(rows)

		if err != nil {
			return nil,
				apperrors.New(
					apperrors.INTERNAL_FALLBACK,
					apperrors.WithMessage("Failed to scan query row"),
					apperrors.WithCauseDetail(err),
				)
		}

		results = append(results, item)
	}

	if err := rows.Err(); err != nil {
		return nil, translateDBError(err)
	}

	return results, nil
}

func QueryOne[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	sql string,
	args []any,
	scan Scanner[T],
) (T, *apperrors.AppError) {

	var zero T

	results, err := QueryMany(
		ctx,
		db,
		sql,
		args,
		scan,
	)

	if err != nil {
		return zero, err
	}

	if len(results) == 0 {
		return zero,
			apperrors.New(
				plerrs.DB_NOT_FOUND,
				apperrors.WithMessage("Failed to build SQL query"),
				apperrors.WithCauseDetail(err),
			)
	}

	return results[0], nil
}

func BuildSQL(qb sq.Sqlizer) (
	string,
	[]any,
	*apperrors.AppError,
) {

	sql, args, err := qb.ToSql()

	if err != nil {
		return "", nil,
			apperrors.New(
				apperrors.INTERNAL_FALLBACK,
				apperrors.WithMessage("Failed to build SQL query"),
				apperrors.WithCauseDetail(err),
			)
	}

	return sql, args, nil
}
