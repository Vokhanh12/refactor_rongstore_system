package reader

import (
	"context"

	authport "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/port"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	authzrepo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/repository"

	"github.com/google/uuid"
)

type AuthorizationReader struct {
	roleAssignmentRepo authzrepo.RoleAssignmentReader
}

var _ authport.AuthorizationReader = (*AuthorizationReader)(nil)

func NewAuthorizationReader(
	roleAssignmentRepo authzrepo.RoleAssignmentReader,
) *AuthorizationReader {
	return &AuthorizationReader{
		roleAssignmentRepo: roleAssignmentRepo,
	}
}

func (a *AuthorizationReader) GetRoleScopesByUserID(
	ctx context.Context,
	userID uuid.UUID,
) ([]security.RoleScope, error) {

	roleScopeResults, err := a.roleAssignmentRepo.GetRoleScopesByUserID(
		ctx,
		query.GetRoleScopesQuery{
			UserID: userID,
		},
	)
	if err != nil {
		return nil, err
	}

	roleScopes := make([]security.RoleScope, 0, len(roleScopeResults))

	for _, result := range roleScopeResults {
		roleScopes = append(roleScopes, security.RoleScope{
			RoleID:    result.RoleID,
			ScopeID:   result.ScopeID,
			ScopeType: result.ScopeType,
		})
	}

	return roleScopes, nil
}
