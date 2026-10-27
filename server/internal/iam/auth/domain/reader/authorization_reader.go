package port

import (
	"context"

	"github.com/google/uuid"
	sec "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
)

type AuthorizationReader interface {
	ListTokenRoleScopes(
		ctx context.Context,
		userID uuid.UUID,
	) ([]sec.TokenRoleScope, error)
}
