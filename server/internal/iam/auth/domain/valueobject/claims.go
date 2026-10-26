package valueobject

import (
	"uuid"

	"github.com/golang-jwt/jwt/v5"
)

type AccessTokenClaims struct {
	jwt.RegisteredClaims

	TenantID     uuid.UUID
	Roles        []TokenRoleScope
	AuthzVersion int
	TokenType    TokenType
}

type RefreshTokenClaims struct {
	jwt.RegisteredClaims

	TokenType TokenType
}
