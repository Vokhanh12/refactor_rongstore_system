package usecases

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type RolePermissionView struct{}

type RolePermissionViewBatch struct {
	Items []core.Operation[RolePermissionView]
}

type ViewRolePermissionUsecase struct {
	repo re.RolePermissionRepository
}

func NewViewRolePermissionUsecase(repo re.RolePermissionRepository) *ViewRolePermissionUsecase {
	return &ViewRolePermissionUsecase{repo: repo}
}

func (u *ViewRolePermissionUsecase) Execute(
	ctx context.Context,
	batch RolePermissionViewBatch,
) *dtos.ViewResultDTO {
	return nil
}
