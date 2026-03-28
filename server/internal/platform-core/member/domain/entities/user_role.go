package entities

import "time"

type ScopeType string

const (
	ScopeGlobal       ScopeType = "GLOBAL"
	ScopeOrganization ScopeType = "ORGANIZATION"
	ScopeUnit         ScopeType = "UNIT"
)

type UserRole struct {
	UserID string
	RoleID string

	ScopeType ScopeType
	ScopeID   string

	AssignedAt time.Time
	AssignedBy *string
}
