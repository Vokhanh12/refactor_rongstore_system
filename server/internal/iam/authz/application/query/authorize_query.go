package query

import (
	"context"

	pr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type AuthorizeQuery interface {
	ListGrantsByRoleRefs(ctx context.Context, roleRefs []vo.RoleRef) ([]pr.AuthorizationGrant, *aerrs.AppError)
}
