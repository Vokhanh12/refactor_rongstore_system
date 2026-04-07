package entities

import (
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type RoleScopeType string
type RoleAcessScope string

const (
	RoleTypeGobal    RoleScopeType  = "GOBAL"
	RoleTypeTenant   RoleScopeType  = "TENANT"
	RoleTypeUNit     RoleScopeType  = "UNIT"
	AcessScopeGobal  RoleAcessScope = "ALL"
	AcessScopeTenant RoleAcessScope = "OWN"
)

type Role struct {
	id      string
	roleRef vo.RoleRef
	name    string

	roleScopeType   RoleScopeType
	roleAccessScope RoleAcessScope
	level           int
	description     *string

	isSystem bool
	isSuper  bool
	isActive bool
}

func NewRole(
	id string,
	roleRef vo.RoleRef,
	roleScopeType RoleScopeType,
	name string,
	roleAccessScope RoleAcessScope,
	level int,
	description *string,
	isSystem, isSuper, isActive bool,
) (Role, []aerrs.AppErrorDetail) {

	var details []aerrs.AppErrorDetail

	if roleScopeType == RoleTypeGobal && roleRef.ScopeID() != "" {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_INVALID_FORMAT,
			aerrs.WithField("scope_id"),
			aerrs.WithMessageDetail("GLOBAL role must not have scope_id"),
		))
	}

	if (roleScopeType == RoleTypeTenant || roleScopeType == RoleTypeUNit) && roleRef.ScopeID() == "" {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_REQUIRED,
			aerrs.WithField("scope_id"),
			aerrs.WithMessageDetail("scope_id is required for tenant/unit role"),
		))
	}

	if level < 0 {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_OUT_OF_RANGE,
			aerrs.WithField("level"),
		))
	}

	if len(details) > 0 {
		return Role{}, details
	}

	return Role{
		id:              id,
		roleRef:         roleRef,
		roleScopeType:   roleScopeType,
		roleAccessScope: roleAccessScope,
		level:           level,
		description:     description,
		isSystem:        isSystem,
		isSuper:         isSuper,
		isActive:        isActive,
	}, nil
}

func NewRoleFromPersistence(
	id string,
	roleRef vo.RoleRef,
	roleScopeType RoleScopeType,
	name string,
	roleAccessScope RoleAcessScope,
	description *string,
	isSystem, isSuper, isActive bool,
) Role {
	return Role{
		id:              id,
		roleRef:         roleRef,
		roleScopeType:   roleScopeType,
		name:            name,
		roleAccessScope: roleAccessScope,
		description:     description,
		isSystem:        isSystem,
		isSuper:         isSuper,
		isActive:        isActive,
	}
}

func (r Role) IsElevated() bool {
	return r.isSuper
}

func (p Role) RoleRef() vo.RoleRef {
	return p.roleRef
}
