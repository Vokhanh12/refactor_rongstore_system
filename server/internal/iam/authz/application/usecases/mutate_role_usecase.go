package usecases

import (
	"context"

	coremp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/mappers"
	coreuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	c "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/common"
	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
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

func (u *MutateRoleUsecase) Execute(ctx context.Context, batch RoleMutationBatch) dtos.MutateResultDTO {

	results := u.engine.Execute(ctx, batch.Items, coremp.BuildMutateResult)

	return dtos.MutateResultDTO{Items: results}
}

func (u *MutateRoleUsecase) handleCreate(
	ctx context.Context,
	cmd c.CreateRoleCommand,
) (*c.CreateRoleCommandResult, *aerrs.AppError) {

	roleRef, err := vo.NewRoleRef(cmd.ScopeID, cmd.Code)
	if err != nil {
		return nil, err
	}

	scope, err := enu.NewRoleAccessScope(cmd.RoleAccessScope)
	if err != nil {
		return nil, err
	}

	scopeType, err := enu.NewRoleScopeType(cmd.RoleScopeType)
	if err != nil {
		return nil, err
	}

	exists, err := u.repo.ExistsByRoleIdentity(ctx, scopeType, roleRef)
	if err != nil {
		return nil, err
	}

	if exists {

	}

	role, err := en.NewRole(en.NewRoleParams{
		ID:              cmd.ID,
		RoleRef:         roleRef,
		RoleScopeType:   scopeType,
		Name:            cmd.Name,
		RoleAccessScope: scope,
		Level:           cmd.Level,
		Description:     cmd.Description,
		IsSystem:        cmd.IsSystem,
		IsSuper:         cmd.IsSuper,
		IsActive:        cmd.IsActive,
	})
	if err != nil {
		return nil, err
	}

	savedRole, err := u.repo.Create(ctx, role)
	if err != nil {
		return nil, err
	}

	return &c.CreateRoleCommandResult{
		Result: common.RoleResult{
			Id:          savedRole.ID(),
			Name:        savedRole.Name(),
			Description: savedRole.Description(),
			CreatedAt:   savedRole.CreatedAt(),
			UpdatedAt:   savedRole.UpdatedAt(),
		},
	}, nil
}

func (u *MutateRoleUsecase) handleUpdate(ctx context.Context, cmd c.UpdateRoleCommand) (*c.UpdateRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}

func (u *MutateRoleUsecase) handleDelete(ctx context.Context, cmd c.DeleteRoleCommand) (*c.DeleteRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}
