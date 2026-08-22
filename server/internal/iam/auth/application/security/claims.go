package security

import "github.com/golang-jwt/jwt/v5"

const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

type AccessTokenClaims struct {
	jwt.RegisteredClaims

	TenantID     string      `json:"tenant_id"`
	Roles        []RoleScope `json:"roles,omitempty"`
	AuthzVersion int         `json:"authz_version"`
	TokenType    string      `json:"token_type"`
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims

	TokenType string `json:"token_type"`
}

type RoleScope struct {
	RoleID    string `json:"role_id"`
	ScopeID   string `json:"scope_id"`
	ScopeType string `json:"scope_type"`
}
