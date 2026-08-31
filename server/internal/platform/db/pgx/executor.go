package pgx

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Scanner[T any] func(pgx.Rows) (T, error)

func QueryMany[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	query squirrel.Sqlizer,
	scan Scanner[T],
) ([]T, error) {

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make([]T, 0)

	for rows.Next() {

		item, err := scan(rows)
		if err != nil {
			return nil, err
		}

		results = append(results, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func QueryOne[T any](
	ctx context.Context,
	db *pgxpool.Pool,
	query squirrel.Sqlizer,
	scan Scanner[T],
) (T, error) {

	var zero T

	results, err := QueryMany(
		ctx,
		db,
		query,
		scan,
	)

	if err != nil {
		return zero, err
	}

	if len(results) == 0 {
		return zero, err
	}

	return results[0], nil
}
