package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

var _ re.RolePermissionRepository = (*SqlcRolePermissionRepository)(nil)

type SqlcRolePermissionRepository struct {
}

// Create implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) Create(ctx context.Context, rolePermission *entities.RolePermission) (*entities.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}

// Delete implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	panic("unimplemented")
}

// FindAllByRoles implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) FindAllByRoles(ctx context.Context, roles []string) ([]*entities.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}

// Update implements [repositories.RolePermissionRepository].
func (s *SqlcRolePermissionRepository) Update(ctx context.Context, rolePermission *entities.RolePermission) (*entities.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}
