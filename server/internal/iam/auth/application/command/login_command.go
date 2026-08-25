package command

type LoginCommand struct {
	Identifier string
	Password   string
}

type LoginCommandResult struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}
