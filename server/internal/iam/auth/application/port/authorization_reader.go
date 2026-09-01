package port

import (
	"context"

	"github.com/google/uuid"
	sec "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/application/security"
)

type AuthorizationReader interface {
	GetRoleScopesByUserID(
		ctx context.Context,
		userID uuid.UUID,
	) ([]sec.RoleScope, error)
}
