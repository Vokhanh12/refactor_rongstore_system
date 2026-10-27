package valueobject

import (
	"uuid"
)

type AccessTokenData struct {
	TenantID     uuid.UUID
	Roles        []TokenRoleScope
	AuthzVersion int
	TokenType    TokenType
}

type RefreshTokenData struct {
	TokenType TokenType
}
