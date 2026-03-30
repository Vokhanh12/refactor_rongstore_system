package command

import (
	"context"

	repos "github.com/vokhanh12/refactor-rongstore-system/server/iam/authz/repositoies"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

// ==================== Create RolePermission ====================
type CreateRolePermissionCommand struct{}
type CreateRolePermissionCommandResult struct{}

type CreateRolePermissionHandler struct {
	rolePermissionRepository repos.RolePermissionRepository
}

func NewCreateRolePermissionHandler(rpRepo repos.RolePermissionRepository) *CreateRolePermissionHandler {
	return &CreateRolePermissionHandler{rolePermissionRepository: rpRepo}
}

func (h *CreateRolePermissionHandler) Handle(
	ctx context.Context,
	cmd CreateRolePermissionCommand,
) (*CreateRolePermissionCommandResult, *aerrs.AppError) {
	// TODO: implement
	return nil, nil
}

// ==================== Update RolePermission ====================
type UpdateRolePermissionCommand struct{}
type UpdateRolePermissionCommandResult struct{}

type UpdateRolePermissionHandler struct {
	rolePermissionRepository repos.RolePermissionRepository
}

func NewUpdateRolePermissionHandler(rpRepo repos.RolePermissionRepository) *UpdateRolePermissionHandler {
	return &UpdateRolePermissionHandler{rolePermissionRepository: rpRepo}
}

func (h *UpdateRolePermissionHandler) Handle(
	ctx context.Context,
	cmd UpdateRolePermissionCommand,
) (*UpdateRolePermissionCommandResult, *aerrs.AppError) {
	// TODO: implement
	return nil, nil
}

// ==================== Delete RolePermission ====================
type DeleteRolePermissionCommand struct{}
type DeleteRolePermissionCommandResult struct{}

type DeleteRolePermissionHandler struct {
	rolePermissionRepository repos.RolePermissionRepository
}

func NewDeleteRolePermissionHandler(rpRepo repos.RolePermissionRepository) *DeleteRolePermissionHandler {
	return &DeleteRolePermissionHandler{rolePermissionRepository: rpRepo}
}

func (h *DeleteRolePermissionHandler) Handle(
	ctx context.Context,
	cmd DeleteRolePermissionCommand,
) (*DeleteRolePermissionCommandResult, *aerrs.AppError) {
	// TODO: implement
	return nil, nil
}
