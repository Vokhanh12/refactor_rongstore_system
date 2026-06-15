package usecases

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/usecase"
	c "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	repos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type RolePermissionMutation struct {
	Create *c.CreateRolePermissionCommand
	Update *c.UpdateRolePermissionCommand
	Delete *c.DeleteRolePermissionCommand
}

type RolePermissionMutationBatch struct {
	Items []core.Operation[RolePermissionMutation]
}

type MutateRolePermissionUsecase struct {
	repo repos.RolePermissionRepository
}

func NewMutateRolePermissionUsecase(repo repos.RolePermissionRepository) *MutateRolePermissionUsecase {
	return &MutateRolePermissionUsecase{repo: repo}
}

func (u *MutateRolePermissionUsecase) Execute(
	ctx context.Context,
	batch RolePermissionMutationBatch,
) *dtos.MutateResultDTO {
	return nil
}
