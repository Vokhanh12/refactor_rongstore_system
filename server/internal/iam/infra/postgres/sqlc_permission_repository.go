package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

var _ re.PermissionRepository = (*SqlcPermissionRepository)(nil)

type SqlcPermissionRepository struct {
	queries *db.Queries
	dberr   dberr.DBError
}

// Create implements [repositories.PermissionRepository].
func (s *SqlcPermissionRepository) Create(ctx context.Context, role *entities.Role) (*entities.Role, *apperrors.AppError) {
	panic("unimplemented")
}

// Delete implements [repositories.PermissionRepository].
func (s *SqlcPermissionRepository) Delete(ctx context.Context, id uuid.UUID) *apperrors.AppError {
	panic("unimplemented")
}

// Update implements [repositories.PermissionRepository].
func (s *SqlcPermissionRepository) Update(ctx context.Context, role *entities.Role) (*entities.Role, *apperrors.AppError) {
	panic("unimplemented")
}
