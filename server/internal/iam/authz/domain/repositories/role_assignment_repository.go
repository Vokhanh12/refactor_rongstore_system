package repositories

import (
	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type RoleAssignmentRepository interface {
	Create(roleAssignment *entities.RoleAssignment) (*entities.RoleAssignment, *aerrs.AppError)
	Update(roleAssignment *entities.RoleAssignment) (*entities.RoleAssignment, *aerrs.AppError)
	Delete(id uuid.UUID) *aerrs.AppError
}
