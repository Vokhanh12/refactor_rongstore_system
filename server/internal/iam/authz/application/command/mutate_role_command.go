package command

import (
	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/common"
)

type CreateRoleCommand struct {
	ID              uuid.UUID
	scope_id        uuid.UUID
	RoleScopeType   string
	Code            string
	Name            string
	Description     string
	RoleAccessScope string
	level           unit8
	IsSystem        bool
	IsActive        bool
	IsSuper         bool
}
type CreateRoleCommandResult struct {
	result common.RoleResult
}

type UpdateRoleCommand struct {
	ID              uuid.UUID
	scope_id        uuid.UUID
	RoleScopeType   string
	Code            string
	Name            string
	Description     string
	RoleAccessScope string
	level           unit8
	IsSystem        bool
	IsActive        bool
	IsSuper         bool
}

type UpdateRoleCommandResult struct{}

type DeleteRoleCommand struct{}
type DeleteRoleCommandResult struct{}
