package security

import aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

type Claims struct {
	Subject   string `json:"sub"`
	Issuer    string `json:"iss,omitempty"`
	Audience  string `json:"aud,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	JTI       string `json:"jti,omitempty"`
}

type AccessTokenClaims struct {
	Claims

	TenantID     string      `json:"tenant_id"`
	Roles        []RoleScope `json:"roles,omitempty"`
	AuthzVersion int         `json:"authz_version"`
}

type RefreshTokenClaims struct {
	Claims

	TokenType string `json:"token_type"`
}

type RoleScope struct {
	RoleID    string `json:"role_id"`
	ScopeID   string `json:"scope_id"`
	ScopeType string `json:"scope_type"`
}

type TokenDecoder interface {
	DecodeAccessToken(
		encoded string,
	) (*AccessTokenClaims, *aerrs.AppError)

	DecodeRefreshToken(
		encoded string,
	) (*RefreshTokenClaims, *aerrs.AppError)
}

type TokenSigner interface {
	SignAccessToken(
		claims AccessTokenClaims,
	) (string, *aerrs.AppError)

	SignRefreshToken(
		claims RefreshTokenClaims,
	) (string, *aerrs.AppError)
}
