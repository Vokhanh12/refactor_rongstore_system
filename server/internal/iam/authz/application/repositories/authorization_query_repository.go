package query

import (
	"context"

	pr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type AuthorizationQueryRepository interface {
	ListGrantsByRoleKeys(ctx context.Context, RoleKeys []vo.RoleKey) ([]pr.AuthorizationGrant, *aerrs.AppError)
}
