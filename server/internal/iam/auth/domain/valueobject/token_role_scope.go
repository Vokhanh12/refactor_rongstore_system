package valueobject

import "uuid"

type TokenRoleScope struct {
	RoleID    uuid.UUID  `json:"role_id"`
	ScopeID   *uuid.UUID `json:"scope_id"`
	ScopeType string     `json:"scope_type"`
}
