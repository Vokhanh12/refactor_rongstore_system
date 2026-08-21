package command

type AuthorizeCommand struct {
	UserID  string
	ScopeID string

	RoleKeyStrs []string

	Resource   string
	Action     string
	ResourceID string
}

type AuthorizeCommandResult struct {
	Allowed bool
}
