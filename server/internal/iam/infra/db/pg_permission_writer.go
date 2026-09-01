package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	repos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ repos.PermissionWriter = (*PermissionWriter)(nil)

type PermissionWriter struct {
	dba *pg.DbAdapter
}

// Create implements [repositories.PermissionWriter].
func (s *PermissionWriter) Create(ctx context.Context, role *entities.Role) (*entities.Role, *apperrors.AppError) {
	panic("unimplemented")
}

// Delete implements [repositories.PermissionWriter].
func (s *PermissionWriter) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	panic("unimplemented")
}

// Update implements [repositories.PermissionWriter].
func (s *PermissionWriter) Update(ctx context.Context, role *entities.Role) (*entities.Role, *apperrors.AppError) {
	panic("unimplemented")
}
