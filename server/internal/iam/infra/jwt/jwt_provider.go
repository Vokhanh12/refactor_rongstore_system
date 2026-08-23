package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/serializer"
	authsur "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ authsur.TokenParser = (*JWTProvider)(nil)
var _ authsur.TokenSigner = (*JWTProvider)(nil)

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

// ParseAccessTokenClaims implements [security.TokenParser].
func (j *JWTProvider) ParseAccessTokenClaims(payload string) (authsur.AccessTokenClaims, *aerr.AppError) {
	var claims authsur.AccessTokenClaims

	if err := serializer.Unmarshal([]byte(payload), &claims); err != nil {
		return authsur.AccessTokenClaims{}, err
	}

	return claims, nil
}

func (j *JWTProvider) SignAccessToken(
	userID string,
	tenantID string,
	roles []authsur.RoleScope,
	authzVersion int,
) (string, *aerr.AppError) {

	now := time.Now()

	claims := authsur.AccessTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID,
			Issuer:    j.issuer,
			Audience:  []string{j.audience},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.accessTTL)),
			ID:        uuid.NewString(),
		},

		TenantID:     tenantID,
		Roles:        roles,
		AuthzVersion: authzVersion,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signed, err := token.SignedString(j.secret)
	if err != nil {
		return "", aerr.New(
			errs.JWT_SIGN_FAILED,
			aerr.WithCauseDetail(err),
		)
	}

	return signed, nil
}

// SignRefreshToken implements [security.TokenSigner].
func (j *JWTProvider) SignRefreshToken(claims authsur.RefreshTokenClaims) (string, *aerr.AppError) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signed, err := token.SignedString(j.secret)
	if err != nil {
		return "", aerr.New(
			errs.JWT_SIGN_FAILED,
			aerr.WithCauseDetail(err),
		)
	}

	return signed, nil
}
