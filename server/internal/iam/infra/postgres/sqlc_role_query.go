package postgres

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	p "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/postgres/fields"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/querydsl"

	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ q.RoleQuery = (*SqlcRoleQuery)(nil)

type SqlcRoleQuery struct {
	dba     *pg.DbAdapter
	builder *querydsl.Builder
}

// Count implements [query.RoleQuery].
func (s *SqlcRoleQuery) Count(ctx context.Context, q q.CountRoleQuery) (q.CountRoleQueryResult, *apperrors.AppError) {
	panic("unimplemented")
}

// Exists implements [query.RoleQuery].
func (s *SqlcRoleQuery) Exists(ctx context.Context, q q.ExistsRoleQuery) (q.ExistsRoleQueryResult, *apperrors.AppError) {
	panic("unimplemented")
}

func NewSqlcRoleQuery(
	dba *pg.DbAdapter,
) q.RoleQuery {

	return &SqlcRoleQuery{
		dba:     dba,
		builder: querydsl.NewBuilder(fields.RoleFields),
	}
}

func (s *SqlcRoleQuery) Search(ctx context.Context, query q.SearchRoleQuery) (
	q.SearchRoleQueryResult,
	*apperrors.AppError,
) {

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

	qb = s.builder.ApplySearch(
		qb,
		query.Criteria.Keyword,
	)

	qb = s.builder.ApplyFilters(
		qb,
		query.Criteria.Filters,
	)

	qb = s.builder.ApplySorts(
		qb,
		query.Criteria.Sorts,
	)

	qb = s.builder.ApplyPagination(
		qb,
		query.Criteria.Pagination,
	)

	sql, args, err := qb.ToSql()

	if err != nil {
		return q.SearchRoleQueryResult{}, apperrors.Internal("BUILD_QUERY_ERROR")
	}

	rows, err := s.dba.P.Query(ctx, sql, args...)

	if err != nil {
		return q.SearchRoleQueryResult{}, apperrors.Internal("QUERY_ROLE_ERROR")
	}

	defer rows.Close()

	results := make([]p.RoleView, 0)

	for rows.Next() {

		var item p.RoleView

		err := rows.Scan(
			&item.ID,
			&item.Code,
			&item.ScopeID,
			&item.Name,
			&item.ScopeType,
			&item.AccessScope,
			&item.Level,
			&item.Description,
			&item.IsSystem,
			&item.IsSuper,
			&item.IsActive,
			&item.CreatedAt,
			&item.UpdatedAt,
		)

		if err != nil {
			return q.SearchRoleQueryResult{},
				apperrors.Internal("SCAN_ROLE_ERROR")
		}

		results = append(
			results,
			item,
		)
	}

	return q.SearchRoleQueryResult{
		Results: results,
	}, nil
}

// Export implements [query.RoleQuery].
func (s *SqlcRoleQuery) Export(ctx context.Context, q q.ExportRoleQuery) (q.ExportRoleQueryResult, *apperrors.AppError) {
	panic("unimplemented")
}

// Get implements [query.RoleQuery].
func (s *SqlcRoleQuery) Get(ctx context.Context, q q.GetRoleQuery) (q.GetRoleQueryResult, *apperrors.AppError) {
	panic("unimplemented")
}

// List implements [query.RoleQuery].
func (s *SqlcRoleQuery) List(ctx context.Context, q q.ListRoleQuery) (q.ListRoleQueryResult, *apperrors.AppError) {
	panic("unimplemented")
}
