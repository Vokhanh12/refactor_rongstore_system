package usecases

import (
	"context"

	coremap "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/mappers"
	coreuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	c "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type RoleMutation struct {
	Create *c.CreateRoleCommand
	Update *c.UpdateRoleCommand
	Delete *c.DeleteRoleCommand
}

type RoleMutationBatch struct {
	Items []coreuc.Operation[RoleMutation]
}

type MutateRoleUsecase struct {
	repo   re.RoleRepository
	engine *coreuc.MutateEngine[RoleMutation]
}

func NewMutateRoleUsecase(repo re.RoleRepository) *MutateRoleUsecase {

	u := &MutateRoleUsecase{
		repo: repo,
	}

	handlers := []coreuc.Handler[RoleMutation]{
		{
			Cond: func(p RoleMutation) bool { return p.Create != nil },
			Exec: func(ctx context.Context, p RoleMutation) (any, *aerrs.AppError) {
				return u.handleCreate(ctx, *p.Create)
			},
		},
		{
			Cond: func(p RoleMutation) bool { return p.Update != nil },
			Exec: func(ctx context.Context, p RoleMutation) (any, *aerrs.AppError) {
				return u.handleUpdate(ctx, *p.Update)
			},
		},
		{
			Cond: func(p RoleMutation) bool { return p.Delete != nil },
			Exec: func(ctx context.Context, p RoleMutation) (any, *aerrs.AppError) {
				return u.handleDelete(ctx, *p.Delete)
			},
		},
	}

	u.engine = coreuc.NewMutateEngine(handlers)

	return u
}

func (u *MutateRoleUsecase) Execute(ctx context.Context, batch RoleMutationBatch) *dtos.MutateResultDTO {

	results := u.engine.Execute(ctx, batch.Items, coremap.BuildMutateResult)

	return &dtos.MutateResultDTO{Items: results}
}

func (u *MutateRoleUsecase) handleCreate(ctx context.Context, cmd c.CreateRoleCommand) (*c.CreateRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}

func (u *MutateRoleUsecase) handleUpdate(ctx context.Context, cmd c.UpdateRoleCommand) (*c.UpdateRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}

func (u *MutateRoleUsecase) handleDelete(ctx context.Context, cmd c.DeleteRoleCommand) (*c.DeleteRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}
