package entities

import vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"

type RoleType string

const (
	RoleTypePlatform     RoleType = "PLATFORM"
	RoleTypeOrganization RoleType = "ORGANIZATION"
	RoleTypeUnit         RoleType = "UNIT"
)

type ScopeType string

const (
	ScopeTypeGobal  ScopeType = "GOBAL"
	ScopeTypeTenant ScopeType = "TENANT"
	ScopeTypeUnit   ScopeType = "UNIT"
	ScopeTypeOwn    ScopeType = "OWN"
)

type Role struct {
	id        string
	roleRef   vo.RoleRef
	scopeType ScopeType
	name      string

	roleType    RoleType
	description *string

	isSystem bool
	isSuper  bool
	isActive bool

	permissions []Permission
}

func (r Role) IsElevated() bool {
	return r.isSuper
}

func (p Role) Key() string {
	return p.roleRef.Role() + ":" + p.Action
}
