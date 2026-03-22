package command

type AuthorizeCommand struct {
	UserID        string
	TenantID      string
	RoleCodes     []string
	ResourceCheck string
	ActionCheck   string
	ResourceID    string
}

type AuthorizeCommandResult struct {
	Allowed bool
}
