package postgres

import (
	"context"

	repos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ repos.RolePermissionQueryRepository = (*SqlcRolePermissionQuery)(nil)

type SqlcRolePermissionQuery struct {
	dba *pg.DbAdapter
}

func NewSqlcRolePermissionQuery(dba *pg.DbAdapter) repos.RolePermissionQueryRepository {
	return &SqlcRolePermissionQuery{dba: dba}
}

// Search implements [query.ViewRolePermissionQuery].
func (s *SqlcRolePermissionQuery) Search(ctx context.Context, a any) (any, *apperrors.AppError) {
	panic("unimplemented")
}
