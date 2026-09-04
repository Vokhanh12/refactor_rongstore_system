package security

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

const (
	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

type AccessTokenClaims struct {
	jwt.RegisteredClaims

	TenantID     uuid.UUID   `json:"tenant_id"`
	Roles        []TokenRole `json:"roles,omitempty"`
	AuthzVersion int         `json:"authz_version"`
	TokenType    string      `json:"token_type"`
}

type TokenRole struct {
	RoleID    uuid.UUID  `json:"role_id"`
	ScopeID   *uuid.UUID `json:"scope_id"`
	ScopeType string     `json:"scope_type"`
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims

	TokenType string `json:"token_type"`
}
