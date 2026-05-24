package postgres

import (
	"context"

	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ q.RoleQuery = (*SqlcRoleQuery)(nil)

type SqlcRoleQuery struct {
	queries *db.Queries
	dberr   dberr.DBError
}

func NewSqlcRoleQuery(queries *db.Queries) q.RoleQuery {
	return &SqlcRoleQuery{queries: queries}
}

// Count implements [query.RoleQuery].
func (s *SqlcRoleQuery) Count(ctx context.Context, q q.CountRoleQuery) (q.CountRoleQueryResult, *apperrors.AppError) {
	panic("unimplemented")
}

// Exists implements [query.RoleQuery].
func (s *SqlcRoleQuery) Exists(ctx context.Context, q q.ExistsRoleQuery) (q.ExistsRoleQueryResult, *apperrors.AppError) {
	panic("unimplemented")
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

// Search implements [query.RoleQuery].
func (s *SqlcRoleQuery) Search(ctx context.Context, q q.SearchRoleQuery) (q.SearchRoleQueryResult, *apperrors.AppError) {
	panic("unimplemented")
}
