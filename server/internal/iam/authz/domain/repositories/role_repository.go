package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type RoleRepository interface {
	Create(ctx context.Context, role *entities.Role) (*entities.Role, *aerrs.AppError)
	Update(ctx context.Context, role *entities.Role) (*entities.Role, *aerrs.AppError)
	Delete(ctx context.Context, id uuid.UUID) *aerrs.AppError
}
