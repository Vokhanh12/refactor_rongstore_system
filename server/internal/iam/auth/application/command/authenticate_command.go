package command

type AuthenticateCommand struct {
	Token string
}

type AuthenticateCommandResult struct {
	Allowed bool
}
