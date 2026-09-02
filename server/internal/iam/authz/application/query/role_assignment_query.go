package query

import (
	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enum"
)

type GetRoleScopesByUserIDQueryResult struct {
	RoleID    uuid.UUID
	ScopeID   *uuid.UUID
	ScopeType enum.RoleScopeType
}
