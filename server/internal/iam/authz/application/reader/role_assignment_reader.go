package reader

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
)

type RoleAssignmentReader interface {
	GetRoleScopesByUserID(ctx context.Context, userID uuid.UUID) ([]query.GetRoleScopesByUserIDQueryResult, error)
}
