package command

type AuthenticateCommand struct {
	Token string
}

type AuthenticateCommandResult struct {
	UserID      string
	RoleKeyStrs []string
}
