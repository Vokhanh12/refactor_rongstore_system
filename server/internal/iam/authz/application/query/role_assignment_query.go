package query

import "github.com/google/uuid"

type GetRoleScopesQuery struct {
	UserID uuid.UUID
}
type GetRoleScopesQueryResult struct {
	RoleID    uuid.UUID
	ScopeID   uuid.UUID
	ScopeType string
}
