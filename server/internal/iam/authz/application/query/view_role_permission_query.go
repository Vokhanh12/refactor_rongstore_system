package command

import (
	"context"

	repos "github.com/vokhanh12/refactor-rongstore-system/server/iam/authz/repositoies"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

// ==================== Search RolePermission ====================
type SearchRolePermissionQuery struct{}
type SearchRolePermissionQueryResult struct{}

type SearchRolePermissionHandler struct {
	rolePermissionRepository repos.RolePermissionRepository
}

func NewSearchRolePermissionHandler(rpRepo repos.RolePermissionRepository) *SearchRolePermissionHandler {
	return &SearchRolePermissionHandler{rolePermissionRepository: rpRepo}
}

func (h *SearchRolePermissionHandler) Handle(
	ctx context.Context,
	cmd SearchRolePermissionQuery,
) (*SearchRolePermissionQueryResult, *aerrs.AppError) {
	// TODO: implement
	return nil, nil
}
