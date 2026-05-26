package postgres

import (
	"context"

	"github.com/google/uuid"
	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ re.RolePermissionRepository = (*SqlcRolePermissionRepository)(nil)

type SqlcRolePermissionRepository struct {
	dba *pg.DbAdapter
}

func NewSqlcRolePermissionRepository(dba *pg.DbAdapter) repositories.RolePermissionRepository {
	return &SqlcRolePermissionRepository{dba: dba}
}

// Create implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) Create(ctx context.Context, rolePermission *en.RolePermission) (*en.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}

// Delete implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	panic("unimplemented")
}

// Update implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) Update(ctx context.Context, rolePermission *en.RolePermission) (*en.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}
