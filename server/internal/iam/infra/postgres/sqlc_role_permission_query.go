package postgres

import (
	"context"

	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

var _ q.ViewRolePermissionQuery = (*SqlcRolePermissionQuery)(nil)

type SqlcRolePermissionQuery struct {
	queries *db.Queries
	dberr   dberr.DBError
}

func NewSqlcRolePermissionQuery(queries *db.Queries) q.ViewRolePermissionQuery {
	return &SqlcRolePermissionQuery{queries: queries}
}

// Search implements [query.ViewRolePermissionQuery].
func (s *SqlcRolePermissionQuery) Search(ctx context.Context, a any) (any, *apperrors.AppError) {
	panic("unimplemented")
}
