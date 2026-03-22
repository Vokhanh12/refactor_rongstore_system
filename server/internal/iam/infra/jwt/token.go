package jwt

import "time"

type TokenService struct {
	signer *Signer
}

func NewTokenService(secret string, issuer string, ttl time.Duration) *TokenService {
	return &TokenService{
		signer: NewSigner(secret, issuer, ttl),
	}
}

func (t *TokenService) IssueToken(userID string, orgID string, unitID string, roles []string) (string, error) {
	return t.signer.Sign(userID, orgID, unitID, roles)
}
