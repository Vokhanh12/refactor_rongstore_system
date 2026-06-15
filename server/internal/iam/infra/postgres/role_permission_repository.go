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

var _ re.RolePermissionRepository = (*RolePermissionRepository)(nil)

type RolePermissionRepository struct {
	dba *pg.DbAdapter
}

func NewRolePermissionRepository(dba *pg.DbAdapter) repositories.RolePermissionRepository {
	return &RolePermissionRepository{dba: dba}
}

// Create implements [repositories.RolePermissionRepository].
func (s *RolePermissionRepository) Create(ctx context.Context, rolePermission *en.RolePermission) (*en.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}

// Delete implements [repositories.RolePermissionRepository].
func (s *RolePermissionRepository) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	panic("unimplemented")
}

// Update implements [repositories.RolePermissionRepository].
func (s *RolePermissionRepository) Update(ctx context.Context, rolePermission *en.RolePermission) (*en.RolePermission, *apperrors.AppError) {
	panic("unimplemented")
}
