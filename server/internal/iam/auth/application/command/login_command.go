package command

type LoginCommand struct {
	Identifier string
	Password   string

	DeviceID  string
	TenantID  string
	Ip        string
	UserAgent string
}

type LoginCommandResult struct {
	AccessToken      string
	RefreshToken     string
	TokenType        string
	ExpiresIn        int64
	RefreshExpiresIn int64
}
