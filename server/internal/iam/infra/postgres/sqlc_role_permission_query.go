package postgres

import (
	"context"

	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ q.RolePermissionQuery = (*SqlcRolePermissionQuery)(nil)

type SqlcRolePermissionQuery struct {
	dba *pg.DbAdapter
}

func NewSqlcRolePermissionQuery(dba *pg.DbAdapter) q.RolePermissionQuery {
	return &SqlcRolePermissionQuery{dba: dba}
}

// Search implements [query.ViewRolePermissionQuery].
func (s *SqlcRolePermissionQuery) Search(ctx context.Context, a any) (any, *apperrors.AppError) {
	panic("unimplemented")
}
