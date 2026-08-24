package command

type LoginCommand struct {
	username string
	password string
}

type LoginCommandResult struct {
	token string
}
