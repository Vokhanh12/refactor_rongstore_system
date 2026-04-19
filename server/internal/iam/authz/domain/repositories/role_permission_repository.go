package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RolePermissionRepository interface {
	Create(ctx context.Context, rolePermission *entities.RolePermission) (*entities.RolePermission, *aerrs.AppError)
	Update(ctx context.Context, rolePermission *entities.RolePermission) (*entities.RolePermission, *aerrs.AppError)
	Delete(ctx context.Context, id uuid.UUID) *aerrs.AppError
	FindAllByRoleRefs(ctx context.Context, roleRefs []vo.RoleRef) ([]*entities.RolePermission, *aerrs.AppError)
}
