package command

type AuthorizeCommand struct {
	UserID      string
	TenantID    string
	RoleKeyStrs []string

	Resource   string
	Action     string
	ResourceID string
}

type AuthorizeCommandResult struct {
	Allowed bool
}
