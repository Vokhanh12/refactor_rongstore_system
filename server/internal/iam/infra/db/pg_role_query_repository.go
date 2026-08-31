package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	repo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/repository"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/db/fields"
	srs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/db/scanrows"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/pgx"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/querydsl"
	sqlc "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"

	sq "github.com/Masterminds/squirrel"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
)

var _ repo.RoleQueryRepository = (*PgRoleQueryRepository)(nil)

type PgRoleQueryRepository struct {
	queries *sqlc.Queries
	pool    *pgxpool.Pool
	builder *querydsl.Builder
}

func NewRoleQueryRepository(q *sqlc.Queries, p *pgxpool.Pool, b *querydsl.Builder) repo.RoleQueryRepository {
	return &PgRoleQueryRepository{
		queries: q,
		pool:    p,
		builder: querydsl.NewBuilder(fields.RoleFields),
	}
}

// Export implements [query.RoleQueryRepository].
func (p *PgRoleQueryRepository) Export(ctx context.Context, q query.ExportRoleQuery) (query.ExportRoleQueryResult, error) {
	panic("unimplemented")
}

// GetById implements [query.RoleQueryRepository].
func (p *PgRoleQueryRepository) GetById(ctx context.Context, q query.GetRoleQuery) (query.GetRoleQueryResult, error) {
	panic("unimplemented")
}

// Search implements [query.RoleQueryRepository].
func (p *PgRoleQueryRepository) Search(ctx context.Context, q query.SearchRoleQuery) (query.SearchRoleQueryResult, error) {

	qb := sq.
		Select(
			"r.id",
			"r.code",
			"r.scope_id",
			"r.name",
			"r.role_scope_type",
			"r.role_access_scope",
			"r.level",
			"r.description",
			"r.is_system",
			"r.is_super",
			"r.is_active",
			"r.created_at",
			"r.updated_at",
		).
		From("roles r").
		PlaceholderFormat(sq.Dollar)

	qb = p.builder.ApplySearch(
		qb,
		q.Criteria.Keyword,
	)

	qb = p.builder.ApplyFilters(
		qb,
		q.Criteria.Filters,
	)

	qb = p.builder.ApplySorts(
		qb,
		q.Criteria.Sorts,
	)

	qb = p.builder.ApplyPagination(
		qb,
		q.Criteria.Pagination,
	)

	results, err := pgx.QueryMany(
		ctx,
		p.pool,
		qb,
		srs.ScanRoleView,
	)

	if err != nil {
		return query.SearchRoleQueryResult{}, err
	}

	return query.SearchRoleQueryResult{
		Items: results,
		Total: len(results),
	}, nil
}
