package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	repos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ repos.PermissionCommandRepository = (*PermissionCommandRepository)(nil)

type PermissionCommandRepository struct {
	dba *pg.DbAdapter
}

// Create implements [repositories.PermissionCommandRepository].
func (s *PermissionCommandRepository) Create(ctx context.Context, role *entities.Role) (*entities.Role, *apperrors.AppError) {
	panic("unimplemented")
}

// Delete implements [repositories.PermissionCommandRepository].
func (s *PermissionCommandRepository) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	panic("unimplemented")
}

// Update implements [repositories.PermissionCommandRepository].
func (s *PermissionCommandRepository) Update(ctx context.Context, role *entities.Role) (*entities.Role, *apperrors.AppError) {
	panic("unimplemented")
}
