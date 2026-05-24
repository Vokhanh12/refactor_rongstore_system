package query

import (
	"context"

	pr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RoleQuery interface {
	Search(ctx context.Context, q SearchRoleQuery) (SearchRoleQueryResult, *aerrs.AppError)
	List(ctx context.Context, q ListRoleQuery) (ListRoleQueryResult, *aerrs.AppError)
	Export(ctx context.Context, q ExportRoleQuery) (ExportRoleQueryResult, *aerrs.AppError)
	Get(ctx context.Context, q GetRoleQuery) (GetRoleQueryResult, *aerrs.AppError)
	Count(ctx context.Context, q CountRoleQuery) (CountRoleQueryResult, *aerrs.AppError)
	Exists(ctx context.Context, q ExistsRoleQuery) (ExistsRoleQueryResult, *aerrs.AppError)
}

type SearchRoleQuery struct{}
type SearchRoleQueryResult struct {
	Items []pr.RoleView
	Total int64
}

type ListRoleQuery struct{}
type ListRoleQueryResult struct{}

type GetRoleQuery struct{}
type GetRoleQueryResult struct{}

type ExportRoleQuery struct{}
type ExportRoleQueryResult struct{}

type CountRoleQuery struct{}
type CountRoleQueryResult struct{}

type ExistsRoleQuery struct{}
type ExistsRoleQueryResult struct{}

type ListRoleByRefQuery struct{}
type ListRoleByRefQueryResult struct{}
