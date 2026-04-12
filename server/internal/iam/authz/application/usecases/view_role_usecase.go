package usecases

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	common "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type RoleView struct{}

type RoleViewBatch struct {
	Items []core.Operation[RoleView]
}

type ViewRoleUsecase struct {
	repo re.RoleRepository
}

func NewViewRoleUsecase(repo re.RoleRepository) *ViewRoleUsecase {
	return &ViewRoleUsecase{repo: repo}
}

func (u *ViewRoleUsecase) Execute(
	ctx context.Context,
	batch RoleViewBatch,
) *common.ViewResult {
	return nil
}
