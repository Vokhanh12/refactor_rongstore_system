package command

import (
	"context"

	repos "github.com/vokhanh12/refactor-rongstore-system/server/iam/authz/repositoies"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

// ==================== Create Role ====================
type CreateRoleCommand struct{}
type CreateRoleCommandResult struct{}

type CreateRoleHandler struct {
	roleRepository repos.RoleRepository
}

func NewCreateRoleHandler(rpRepo repos.RoleRepository) *CreateRoleHandler {
	return &CreateRoleHandler{roleRepository: rpRepo}
}

func (h *CreateRoleHandler) Handle(ctx context.Context, cmd CreateRoleCommand) (*CreateRoleCommandResult, *aerrs.AppError) {

	return nil, nil
}

// ==================== Update Role ====================
type UpdateRoleCommand struct{}
type UpdateRoleCommandResult struct{}

type UpdateRoleHandler struct {
	roleRepository repos.RoleRepository
}

func NewUpdateRoleHandler(rpRepo repos.RoleRepository) *UpdateRoleHandler {
	return &UpdateRoleHandler{roleRepository: rpRepo}
}

func (h *UpdateRoleHandler) Handle(ctx context.Context, cmd UpdateRoleCommand) (*UpdateRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}

// ==================== Delete Role ====================
type DeleteRoleCommand struct{}
type DeleteRoleCommandResult struct{}

type DeleteRoleHandler struct {
	roleRepository repos.RoleRepository
}

func NewDeleteRoleHandler(rpRepo repos.RoleRepository) *DeleteRoleHandler {
	return &DeleteRoleHandler{roleRepository: rpRepo}
}

func (h *DeleteRoleHandler) Handle(ctx context.Context, cmd DeleteRoleCommand) (*DeleteRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}
