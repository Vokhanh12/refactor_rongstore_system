package query

import (
	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enum"
)

type ListRoleScopesQueryResult struct {
	RoleID    uuid.UUID
	ScopeID   *uuid.UUID
	ScopeType enum.RoleScopeType
}
