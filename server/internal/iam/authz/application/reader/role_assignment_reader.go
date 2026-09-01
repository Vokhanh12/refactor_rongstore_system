package reader

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
)

type RoleAssignmentReader interface {
	GetRoleScopesByUserID(ctx context.Context, query query.GetRoleScopesQuery) ([]query.GetRoleScopesQueryResult, error)
}
