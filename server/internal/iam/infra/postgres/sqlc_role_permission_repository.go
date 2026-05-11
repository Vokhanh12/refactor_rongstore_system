package postgres

import (
	"context"

	"github.com/google/uuid"
	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ re.RolePermissionRepository = (*SqlcRolePermissionRepository)(nil)

type SqlcRolePermissionRepository struct {
	queries *db.Queries
	dberr   dberr.DBError
}

func NewSqlcRolePermissionRepository(queries *db.Queries) repositories.RolePermissionRepository {
	return &SqlcRolePermissionRepository{queries: queries}
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
