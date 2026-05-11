package command

type AuthorizationGrant struct {
	RoleRef vo.RoleRef

	IsElevated bool

	Resource string
	Action   string
}

type AuthorizeCommand struct {
	UserID   string
	TenantID string
	Roles    []string

	Resource   string
	Action     string
	ResourceID string
}

type AuthorizeCommandResult struct {
	Allowed bool
}
