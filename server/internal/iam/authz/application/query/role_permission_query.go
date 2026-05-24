package query

import (
	"context"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RolePermissionQuery interface {
	Search(ctx context.Context, a any) (any, *aerrs.AppError)
}

type SearchRolePermissionQuery struct{}
type SearchRolePermissionQueryResult struct{}

type GetRolePermissionQuery struct{}
type GetRolePermissionQueryResult struct{}

type ListRolePermissionQuery struct{}
type ListRolePermissionQueryResult struct{}
