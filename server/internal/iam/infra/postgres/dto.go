package postgres

import "github.com/google/uuid"

type roleRefDTO struct {
	RoleCode string     `json:"role_code"`
	ScopeID  *uuid.UUID `json:"scope_id"`
}
