package command

type AuthenticateCommand struct {
	AccessToken string
}

type AuthenticateCommandResult struct {
	Allowed bool
}
