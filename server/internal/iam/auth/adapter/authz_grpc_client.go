package adapter

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	authzrepo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/reader"
)

type AuthzGrpcClient struct {
	roleAssignmentReader authzrepo.RoleAssignmentReader
}

var _ authzrepo.RoleAssignmentReader = (*AuthzGrpcClient)(nil)

func NewAuthzGrpcClient(
	rare authzrepo.RoleAssignmentReader,
) *AuthzGrpcClient {
	return &AuthzGrpcClient{
		roleAssignmentReader: rare,
	}
}

// ListRoleScopes implements [reader.RoleAssignmentReader].
func (a *AuthzGrpcClient) ListRoleScopes(ctx context.Context, userID uuid.UUID) ([]query.ListRoleScopesQueryResult, error) {

	roleScopeResults, err := a.roleAssignmentReader.ListRoleScopes(
		ctx,
		userID,
	)
	if err != nil {
		return nil, err
	}

	roleScopes := make([]ser.TokenRoleScope, 0, len(roleScopeResults))

	for _, result := range roleScopeResults {
		roleScopes = append(roleScopes, ser.TokenRoleScope{
			RoleID:    result.RoleID,
			ScopeID:   result.ScopeID,
			ScopeType: string(result.ScopeType),
		})
	}

	return roleScopes, nil
}
