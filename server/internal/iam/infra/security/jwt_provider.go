package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/serializer"
	as "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/security"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/valueobject"
)

var _ as.TokenParser = (*JWTProvider)(nil)
var _ as.TokenSigner = (*JWTProvider)(nil)

type JWTProvider struct {
	secret     []byte
	issuer     string
	audience   string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJWTProvider(secret []byte, issuer string, audience string, accessTTL time.Duration, refreshTTL time.Duration) *JWTProvider {
	return &JWTProvider{
		secret:     secret,
		issuer:     issuer,
		audience:   audience,
		accessTTL:  accessTTL,
		refreshTTL: refreshTTL,
	}
}

// ParseAccessToken implements [security.TokenParser].
func (j *JWTProvider) ParseAccessToken(payload string) (vo.AccessTokenData, error) {
	var claims vo.AccessTokenData

	if err := serializer.Unmarshal([]byte(payload), &claims); err != nil {
		return vo.AccessTokenData{}, err
	}

	return claims, nil
}

// ParseRefreshToken implements [security.TokenParser].
func (j *JWTProvider) ParseRefreshToken(payload string) (vo.RefreshTokenData, error) {
	panic("unimplemented")
}

func (j *JWTProvider) SignAccessToken(
	userID string,
	roles []vo.TokenRoleScope,
	authzVersion int,
) (string, error) {

	now := time.Now()

	data := vo.AccessTokenData{
		Roles:        roles,
		AuthzVersion: authzVersion,
	}

	claims := accessTokenClaims{
		AccessTokenData: data,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    j.issuer,
			Audience:  []string{j.audience},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.refreshTTL)),
			ID:        uuid.NewString(),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signed, err := token.SignedString(j.secret)
	if err != nil {
		return "", err
	}

	return signed, nil
}

// SignRefreshToken implements [security.TokenSigner].
func (j *JWTProvider) SignRefreshToken(
	userID string,
) (string, error) {
	now := time.Now()

	data := vo.RefreshTokenData{}

	claims := refreshTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    j.issuer,
			Audience:  []string{j.audience},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.refreshTTL)),
			ID:        uuid.NewString(),
		},
		RefreshTokenData: data,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(j.secret)
}
