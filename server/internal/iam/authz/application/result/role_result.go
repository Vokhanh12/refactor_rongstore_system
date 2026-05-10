package mapper

import (
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/common"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
)

func NewRoleResultFromEntity(role *entities.Role) *common.RoleResult {
	if role == nil {
		return nil
	}

	return &common.RoleResult{
		Id:          role.ID().String(),
		Name:        role.Name(),
		Description: *role.Description(),
		CreatedAt:   role.CreatedAt(),
		UpdatedAt:   role.UpdatedAt(),
	}
}
