package entities

type RoleType string

const (
	RoleTypePlatform     RoleType = "PLATFORM"
	RoleTypeOrganization RoleType = "ORGANIZATION"
	RoleTypeUnit         RoleType = "UNIT"
)

type Role struct {
	ID   string
	Code string
	Name string

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
