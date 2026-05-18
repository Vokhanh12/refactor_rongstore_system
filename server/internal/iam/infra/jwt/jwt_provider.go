package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authports "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/domain/ports"
	errs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/errors"
	authzports "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/ports"
	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ authports.TokenGenerator = (*JWTProvider)(nil)
var _ authzports.TokenProvider = (*JWTProvider)(nil)
var _ authports.TokenSigner = (*JWTProvider)(nil)

type JWTProvider struct {
	secret []byte
	issuer string
	ttl    time.Duration
}

func NewJWTProvider(secret []byte, issuer string, ttl time.Duration) *JWTProvider {
	return &JWTProvider{
		secret: secret,
		issuer: issuer,
		ttl:    ttl,
	}
}

func (j *JWTProvider) Sign(userID string, orgID string, unitID string, roles []string) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: userID,
		OrgID:  orgID,
		UnitID: unitID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.secret)
}

// DecorePayload implements [ports.TokenProvider].
func (j *JWTProvider) DecorePayload(encoded string) (*authzports.Payload, *aerr.AppError) {
	if encoded == "" {
		return nil,
			aerr.New(
				errs.JWT_INVALID,
				aerr.WithCauseDetail(
					errors.New("empty jwt payload"),
				),
			)
	}

	raw, err := base64.RawURLEncoding.DecodeString(encoded)

	if err != nil {
		return nil,
			aerr.New(
				errs.JWT_PAYLOAD_INVALID,
				aerr.WithCauseDetail(err),
			)
	}

	var payload authzports.Payload

	if err := json.Unmarshal(raw, &payload); err != nil {

		return nil,
			aerr.New(
				errs.JWT_PAYLOAD_INVALID,
				aerr.WithCauseDetail(err),
			)
	}

	return &payload, nil
}

// GenerateAccessToken implements [services.TokenProvider].
func (j *JWTProvider) GenerateAccessToken(userID string, tenantID string) (string, error) {
	panic("unimplemented")
}

// GenerateRefreshToken implements [services.TokenProvider].
func (j *JWTProvider) GenerateRefreshToken(userID string) (string, error) {
	panic("unimplemented")
}
