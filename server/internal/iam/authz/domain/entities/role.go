package entities

import vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"

type RoleType string

const (
	RoleTypePlatform     RoleType = "PLATFORM"
	RoleTypeOrganization RoleType = "ORGANIZATION"
	RoleTypeUnit         RoleType = "UNIT"
)

type RoleScopeType string

const (
	ScopeTypeGobal  RoleScopeType = "GOBAL"
	ScopeTypeTenant RoleScopeType = "TENANT"
	ScopeTypeUnit   RoleScopeType = "UNIT"
	ScopeTypeOwn    RoleScopeType = "OWN"
)

type Role struct {
	id        string
	roleRef   vo.RoleRef
	scopeType RoleScopeType
	name      string

	roleType    RoleType
	description string

	isSystem bool
	isSuper  bool
	isActive bool
}

func NewRoleFromPersistence(
	id string,
	roleRef vo.RoleRef,
	scopeType RoleScopeType,
	name string,
	roleType RoleType,
	description string,
	isSystem, isSuper, isActive bool,
) Role {
	return Role{
		id:          id,
		roleRef:     roleRef,
		scopeType:   scopeType,
		name:        name,
		roleType:    roleType,
		description: description,
		isSystem:    isSystem,
		isSuper:     isSuper,
		isActive:    isActive,
	}
}

func (r Role) IsElevated() bool {
	return r.isSuper
}

func (p Role) RoleRef() vo.RoleRef {
	return p.roleRef
}
