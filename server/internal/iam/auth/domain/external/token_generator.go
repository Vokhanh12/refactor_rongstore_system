package external

type TokenGenerator interface {
	GenerateAccessToken(userID string, tenantID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
}
