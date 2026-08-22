package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	authsur "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"

	errs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/error"
	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ authsur.TokenDecoder = (*JWTProvider)(nil)
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

// DecodeAccessToken implements [security.TokenDecoder].
func (j *JWTProvider) DecodeAccessToken(encoded string) (*authsur.AccessTokenClaims, *aerr.AppError) {
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

	var token authsur.AccessTokenClaims

	if err := json.Unmarshal(raw, &token); err != nil {

		return nil,
			aerr.New(
				errs.JWT_PAYLOAD_INVALID,
				aerr.WithCauseDetail(err),
			)
	}

	return &token, nil
}

// DecodeRefreshToken implements [security.TokenDecoder].
func (j *JWTProvider) DecodeRefreshToken(
	encoded string,
) (*authsur.RefreshTokenClaims, *aerr.AppError) {

	if encoded == "" {
		return nil, aerr.New(
			errs.JWT_INVALID,
			aerr.WithCauseDetail(
				errors.New("empty jwt"),
			),
		)
	}

	var claims authsur.RefreshTokenClaims

	token, err := jwt.ParseWithClaims(
		encoded,
		&claims,
		func(token *jwt.Token) (any, error) {

			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}

			return j.secret, nil
		},
		jwt.WithIssuer(j.issuer),
		jwt.WithValidMethods([]string{
			jwt.SigningMethodHS256.Alg(),
		}),
	)

	if err != nil {
		return nil, aerr.New(
			errs.JWT_INVALID,
			aerr.WithCauseDetail(err),
		)
	}

	if !token.Valid {
		return nil, aerr.New(
			errs.JWT_INVALID,
			aerr.WithCauseDetail(
				errors.New("invalid jwt"),
			),
		)
	}

	if claims.Subject == "" {
		return nil, aerr.New(
			errs.JWT_INVALID,
			aerr.WithCauseDetail(
				errors.New("missing subject"),
			),
		)
	}

	if claims.TokenType != tokenTypeRefresh {
		return nil, aerr.New(
			errs.JWT_INVALID,
			aerr.WithCauseDetail(
				errors.New("not a refresh token"),
			),
		)
	}

	return &claims, nil
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
