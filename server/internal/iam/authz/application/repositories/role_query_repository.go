package query

import (
	"context"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RoleQueryRepository interface {
	Search(ctx context.Context, q SearchRoleQuery) (SearchRoleQueryResult, *aerrs.AppError)
	Export(ctx context.Context, q ExportRoleQuery) (ExportRoleQueryResult, *aerrs.AppError)
	GetById(ctx context.Context, q GetRoleQuery) (GetRoleQueryResult, *aerrs.AppError)
}
