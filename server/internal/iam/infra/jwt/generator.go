package jwt

import services "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/ports"

var _ services.TokenGenerator = (*JWTGenerator)(nil)

type JWTGenerator struct {
	secret string
}

func NewJWTGenerator(secret string) *JWTGenerator {
	return &JWTGenerator{
		secret: secret,
	}
}

// GenerateAccessToken implements [services.TokenGenerator].
func (j *JWTGenerator) GenerateAccessToken(userID string, tenantID string) (string, error) {
	panic("unimplemented")
}

// GenerateRefreshToken implements [services.TokenGenerator].
func (j *JWTGenerator) GenerateRefreshToken(userID string) (string, error) {
	panic("unimplemented")
}
