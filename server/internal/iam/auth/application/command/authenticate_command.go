package command

type AuthenticateCommand struct {
	Payload string
}

type AuthenticateCommandResult struct {
	Allowed bool
}
