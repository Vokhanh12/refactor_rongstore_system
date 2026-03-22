package repositories

import (
	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RolePermissionRepository interface {
	Create(rolePermission *entities.RolePermission) (*entities.RolePermission, *aerrs.AppError)
	Update(rolePermission *entities.RolePermission) (*entities.RolePermission, *aerrs.AppError)
	Delete(id uuid.UUID) *aerrs.AppError
	FindAllByRoleCode(roleCode string) ([]*entities.RolePermission, *aerrs.AppError)
}
