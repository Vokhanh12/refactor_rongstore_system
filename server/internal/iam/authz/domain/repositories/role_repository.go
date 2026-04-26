package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RoleRepository interface {
	Create(ctx context.Context, role *entities.Role) (*entities.Role, *aerrs.AppError)
	Update(ctx context.Context, role *entities.Role) (*entities.Role, *aerrs.AppError)
	Delete(ctx context.Context, id uuid.UUID) *aerrs.AppError
	FindById(ctx context.Context, id uuid.UUID) (*entities.Role, *aerrs.AppError)
	FindByCode(ctx context.Context, code string) (*entities.Role, *aerrs.AppError)
	Exists(ctx context.Context, roleScopeType enu.RoleScopeType, roleRef vo.RoleRef) (bool, *aerrs.AppError)
}
