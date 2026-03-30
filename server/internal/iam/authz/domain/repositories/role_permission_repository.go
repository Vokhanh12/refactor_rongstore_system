package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type RolePermissionRepository interface {
	Create(ctx context.Context, rolePermission *entities.RolePermission) (*entities.RolePermission, *aerrs.AppError)
	Update(ctx context.Context, rolePermission *entities.RolePermission) (*entities.RolePermission, *aerrs.AppError)
	Delete(ctx context.Context, id uuid.UUID) *aerrs.AppError
	FindAllByRoles(ctx context.Context, roles []string) ([]*entities.RolePermission, *aerrs.AppError)
}
