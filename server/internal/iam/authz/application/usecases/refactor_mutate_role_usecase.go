package usecases

import (
	"context"

	repos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

	c "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	mapper "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/result"
)

const (
	RoleCreate = "role.create"
	RoleUpdate = "role.update"
)

type HandlerFunc func(ctx context.Context, payload any) (any, *aerrs.AppError)

type Dispatcher struct {
	handlers map[string]HandlerFunc
}

func NewDispatcher() *Dispatcher {

	return &Dispatcher{
		handlers: make(map[string]HandlerFunc),
	}

}

func (d *Dispatcher) Register(operation string, handler HandlerFunc) {

	d.handlers[operation] = handler

}

func (d *Dispatcher) Dispatch(
	ctx context.Context,

	operation string,

	payload any,

) (any, *aerrs.AppError) {

	handler, ok := d.handlers[operation]

	if !ok {
		return nil, ErrHandlerNotFound
	}

	return handler(ctx, payload)

}

type MutateRoleUsecase struct {
	repo       repos.RoleRepository
	dispatcher *Dispatcher
}

func NewMutateRoleUsecase(
	repo repos.RoleRepository,
) *MutateRoleUsecase {

	u := &MutateRoleUsecase{
		repo:       repo,
		dispatcher: NewDispatcher(),
	}

	u.dispatcher.Register(
		RoleCreate,
		u.dispatchCreate,
	)

	return u
}

func (u *MutateRoleUsecase) dispatchCreate(
	ctx context.Context,
	payload any,
) (any, *aerrs.AppError) {

	cmd := payload.(c.CreateRoleCommand)

	return u.handleCreate(ctx, cmd)

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

	exists, err := u.command.ExistsRoleByCodeScope(ctx, scopeType, roleKey)
	if err != nil {
		return nil, coreuc.Translate(err)
	}

	if exists {
		return nil, aerrs.New(core.ENTITY_ALREADY_EXISTS)
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

	savedRole, err := u.command.Create(ctx, role)
	if err != nil {
		return nil, coreuc.Translate(err)
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
