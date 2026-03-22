package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Signer struct {
	secret []byte
	issuer string
	ttl    time.Duration
}

func NewSigner(secret string, issuer string, ttl time.Duration) *Signer {
	return &Signer{
		secret: []byte(secret),
		issuer: issuer,
		ttl:    ttl,
	}
}

func (s *Signer) Sign(userID string, orgID string, unitID string, roles []string) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: userID,
		OrgID:  orgID,
		UnitID: unitID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(s.secret)
}
