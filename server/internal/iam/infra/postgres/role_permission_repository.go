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

var _ re.RolePermissionCommandRepository = (*RolePermissionCommandRepository)(nil)

type RolePermissionCommandRepository struct {
	dba *pg.DbAdapter
}

func NewRolePermissionCommandRepository(dba *pg.DbAdapter) repositories.RolePermissionCommandRepository {
	return &RolePermissionCommandRepository{dba: dba}
}

// Create implements [repositories.RolePermissionRepository].
func (s *RolePermissionCommandRepository) Create(ctx context.Context, rolePermission *en.RolePermission) (*en.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}

// Delete implements [repositories.RolePermissionRepository].
func (s *RolePermissionCommandRepository) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	panic("unimplemented")
}

// Update implements [repositories.RolePermissionRepository].
func (s *RolePermissionCommandRepository) Update(ctx context.Context, rolePermission *en.RolePermission) (*en.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}
