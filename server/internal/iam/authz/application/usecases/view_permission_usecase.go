package usecases

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/usecase"
	repos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type PermissionView struct{}

type PermissionViewBatch struct {
	Items []core.Operation[PermissionView]
}

type ViewPermissionUsecase struct {
	repo repos.PermissionQueryRepository
}

func NewViewPermissionUsecase(repo repos.PermissionQueryRepository) *ViewPermissionUsecase {
	return &ViewPermissionUsecase{repo: repo}
}

func (u *ViewPermissionUsecase) Execute(
	ctx context.Context,
	batch PermissionViewBatch,
) *dtos.ViewResultDTO {
	return nil
}
