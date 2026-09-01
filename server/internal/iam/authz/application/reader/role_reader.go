package reader

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
)

type RoleReader interface {
	Search(ctx context.Context, q query.SearchRoleQuery) (query.SearchRoleQueryResult, error)
	Export(ctx context.Context, q query.ExportRoleQuery) (query.ExportRoleQueryResult, error)
	GetById(ctx context.Context, q query.GetRoleQuery) (query.GetRoleQueryResult, error)
}
