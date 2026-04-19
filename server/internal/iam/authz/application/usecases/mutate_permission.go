package usecases

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	c "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type PermissionMutation struct {
	Create *c.CreatePermissionCommand
	Update *c.UpdatePermissionCommand
	Delete *c.DeletePermissionCommand
}

type PermissionMutationBatch struct {
	Items []core.Operation[PermissionMutation]
}

type MutatePermissionUsecase struct {
	repo re.PermissionRepository
}

func NewMutatePermissionUsecase(repo re.PermissionRepository) *MutatePermissionUsecase {
	return &MutatePermissionUsecase{repo: repo}
}

func (u *MutatePermissionUsecase) Execute(
	ctx context.Context,
	batch PermissionMutationBatch,
) *dtos.MutateResultDTO {
	return nil
}
