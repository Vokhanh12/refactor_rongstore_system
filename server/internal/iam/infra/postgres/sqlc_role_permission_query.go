package postgres

import (
	"context"

	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ q.RolePermissionQuery = (*SqlcRolePermissionQuery)(nil)

type SqlcRolePermissionQuery struct {
	queries *db.Queries
	dberr   dberr.DBError
}

func NewSqlcRolePermissionQuery(queries *db.Queries) q.RolePermissionQuery {
	return &SqlcRolePermissionQuery{queries: queries}
}

// Search implements [query.ViewRolePermissionQuery].
func (s *SqlcRolePermissionQuery) Search(ctx context.Context, a any) (any, *apperrors.AppError) {
	panic("unimplemented")
}
