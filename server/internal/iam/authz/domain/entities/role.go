package entities

type RoleType string

const (
	RoleTypePlatform     RoleType = "PLATFORM"
	RoleTypeOrganization RoleType = "ORGANIZATION"
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

	// Aggregate relation (optional, nếu load full)
	Permissions []Permission
}
