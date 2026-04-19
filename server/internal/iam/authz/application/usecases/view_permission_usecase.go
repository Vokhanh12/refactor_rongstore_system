package usecases

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type PermissionView struct{}

type PermissionViewBatch struct {
	Items []core.Operation[PermissionView]
}

type ViewPermissionUsecase struct {
	repo re.PermissionRepository
}

func NewViewPermissionUsecase(repo re.PermissionRepository) *ViewPermissionUsecase {
	return &ViewPermissionUsecase{repo: repo}
}

func (u *ViewPermissionUsecase) Execute(
	ctx context.Context,
	batch PermissionViewBatch,
) *dtos.ViewResultDTO {
	return nil
}
