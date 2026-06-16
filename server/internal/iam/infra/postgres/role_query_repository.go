package postgres

import (
	"context"

	"github.com/google/uuid"
	authzrepos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/postgres/fields"
	srs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/infra/postgres/scanrows"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/querydsl"

	sq "github.com/Masterminds/squirrel"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ authzrepos.RoleQueryRepository = (*RoleQueryRepository)(nil)

type RoleQueryRepository struct {
	dba     *pg.DbAdapter
	builder *querydsl.Builder
}

func NewRoleQueryRepository(
	dba *pg.DbAdapter,
) authzrepos.RoleQueryRepository {

	return &RoleQueryRepository{
		dba:     dba,
		builder: querydsl.NewBuilder(fields.RoleFields),
	}
}

// FindByCode implements [repositories.RoleQueryRepository].
func (s *RoleQueryRepository) FindByCode(ctx context.Context, code string) (*entities.Role, *apperrors.AppError) {
	panic("unimplemented")
}

// FindById implements [repositories.RoleQueryRepository].
func (s *RoleQueryRepository) FindById(ctx context.Context, id uuid.UUID) (*entities.Role, *apperrors.AppError) {
	panic("unimplemented")
}
func (s *RoleQueryRepository) Search(ctx context.Context, query q.SearchRoleQuery) (
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

	sql, args, err := pg.BuildSQL(qb)

	if err != nil {
		return q.SearchRoleQueryResult{}, err
	}

	results, err := pg.QueryMany(
		ctx,
		s.dba.P,
		sql,
		args,
		srs.ScanRoleView,
	)

	if err != nil {
		return q.SearchRoleQueryResult{}, err
	}

	return q.SearchRoleQueryResult{
		Items: results,
		Total: len(results),
	}, nil
}

// Export implements [query.RoleQuery].
func (s *RoleQueryRepository) Export(ctx context.Context, q q.ExportRoleQuery) (q.ExportRoleQueryResult, *apperrors.AppError) {
	panic("unimplemented")
}

// Get implements [query.RoleQuery].
func (s *RoleQueryRepository) Get(ctx context.Context, q q.GetRoleQuery) (q.GetRoleQueryResult, *apperrors.AppError) {
	panic("unimplemented")
}
