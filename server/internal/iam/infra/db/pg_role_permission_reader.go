package postgres

import (
	"context"

	repos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ repos.RolePermissionReader = (*SqlcRolePermissionReader)(nil)

type SqlcRolePermissionReader struct {
	dba *pg.DbAdapter
}

func NewSqlcRolePermissionReader(dba *pg.DbAdapter) repos.RolePermissionReader {
	return &SqlcRolePermissionReader{dba: dba}
}

// Search implements [Reader.ViewRolePermissionReader].
func (s *SqlcRolePermissionReader) Search(ctx context.Context, a any) (any, *apperrors.AppError) {
	panic("unimplemented")
}
