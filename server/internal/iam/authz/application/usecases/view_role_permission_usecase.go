package usecases

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/usecase"
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	repos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type RolePermissionView struct {
	Get *q.GetRoleQuery
}

type RolePermissionViewBatch struct {
	Items []core.Operation[RolePermissionView]
}

type ViewRolePermissionUsecase struct {
	repo repos.RolePermissionQueryRepository
}

func NewViewRolePermissionUsecase(repo repos.RolePermissionQueryRepository) *ViewRolePermissionUsecase {
	return &ViewRolePermissionUsecase{repo: repo}
}

func (u *ViewRolePermissionUsecase) Execute(
	ctx context.Context,
	batch RolePermissionViewBatch,
) *dtos.ViewResultDTO {
	return nil
}
