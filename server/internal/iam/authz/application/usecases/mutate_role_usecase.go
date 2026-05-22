package usecases

import (
	"context"

	coreuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/usecase"
	c "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	mapper "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/result"
	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
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

	results := u.engine.Execute(ctx, batch.Items)

	return dtos.MutateResultDTO{Items: results}
}

func (u *MutateRoleUsecase) handleCreate(
	ctx context.Context,
	cmd c.CreateRoleCommand,
) (*c.CreateRoleCommandResult, *aerrs.AppError) {

	roleKey, err := vo.NewRoleKey(cmd.ScopeID, cmd.Code)
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

	exists, err := u.repo.Exists(ctx, scopeType, roleKey)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, aerrs.New(authzerr.ROLE_ALREADY_EXISTS)
	}

	role, err := en.NewRole(
		en.RolePayload{
			RoleKey:         roleKey,
			RoleScopeType:   scopeType,
			Name:            cmd.Name,
			RoleAccessScope: scope,
			Level:           cmd.Level,
			Description:     cmd.Description,
			IsSystem:        cmd.IsSystem,
			IsSuper:         cmd.IsSuper,
			IsActive:        cmd.IsActive,
		},
	)

	if err != nil {
		return nil, err
	}

	savedRole, err := u.repo.Create(ctx, role)
	if err != nil {
		return nil, err
	}

	return &c.CreateRoleCommandResult{
		Result: mapper.NewRoleResultFromEntity(savedRole),
	}, nil
}

func (u *MutateRoleUsecase) handleUpdate(ctx context.Context, cmd c.UpdateRoleCommand) (*c.UpdateRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}

func (u *MutateRoleUsecase) handleDelete(ctx context.Context, cmd c.DeleteRoleCommand) (*c.DeleteRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}
