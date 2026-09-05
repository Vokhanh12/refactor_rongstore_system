package mapper

import (
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
)

func ToIdentityContext(claims security.AccessTokenClaims) ctxutil.IdentityContext {
	return ctxutil.IdentityContext{
		UserID:       claims.Subject,
		RoleScopes:   ToRoleScopes(claims.Roles),
		AuthzVersion: claims.AuthzVersion,
	}
}

func ToRoleScopes(roles []security.TokenRoleScope) []ctxutil.RoleScope {
	if len(roles) == 0 {
		return nil
	}

	result := make([]ctxutil.RoleScope, 0, len(roles))

	for _, role := range roles {
		result = append(result, ctxutil.RoleScope{
			RoleID:    role.RoleID,
			ScopeID:   role.ScopeID,
			ScopeType: role.ScopeType,
		})
	}

	return result
}
