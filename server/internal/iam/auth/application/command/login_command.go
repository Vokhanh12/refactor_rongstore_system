package command

type LoginCommand struct {
	Identifier string
	Password   string
}

type LoginCommandResult struct {
	AccessToken      string
	RefreshToken     string
	TokenType        string
	ExpiresIn        int64
	RefreshExpiresIn int64
}
