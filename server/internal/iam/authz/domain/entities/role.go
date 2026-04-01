package entities

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
	ID        string
	ScopeID   string
	ScopeType ScopeType
	Code      string
	Name      string

	Type        RoleType
	Description *string

	IsSystem bool
	IsSuper  bool
	IsActive bool

	Permissions []Permission
}

func (r Role) IsElevated() bool {
	return r.IsSuper
}

func (p Permission) Key() string {
	return p.Resource + ":" + p.Action
}
