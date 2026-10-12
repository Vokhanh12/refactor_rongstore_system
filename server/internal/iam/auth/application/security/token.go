package security

type TokenParser interface {
	ParseAccessToken(payload string) (AccessTokenClaims, error)
}

type TokenSigner interface {
	SignAccessToken(
		userID string,
		roles []TokenRoleScope,
		authzVersion int,
	) (string, error)

	SignRefreshToken(
		userID string,
	) (string, error)
}
