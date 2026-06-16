package command

import (
	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/common"
)

type CreateRoleCommand struct {
	ScopeID         *uuid.UUID
	RoleScopeType   string
	Code            string
	Name            string
	Description     *string
	RoleAccessScope string
	Level           int32
	IsSystem        bool
	IsActive        bool
	IsSuper         bool
}

type CreateRoleCommandResult struct {
	Result *common.RoleResult
}

type UpdateRoleCommand struct {
	ID              uuid.UUID
	ScopeID         *uuid.UUID
	RoleScopeType   string
	Code            string
	Name            string
	Description     *string
	RoleAccessScope string
	Level           int32
	IsSystem        bool
	IsActive        bool
	IsSuper         bool
}

type UpdateRoleCommandResult struct {
	Result common.RoleResult
}

type DeleteRoleCommand struct {
	ID uuid.UUID
}
type DeleteRoleCommandResult struct{}

type ExistsRoleQuery struct {
	Code string
}

type ExistsRoleQueryResult struct {
	Exsits bool
}
