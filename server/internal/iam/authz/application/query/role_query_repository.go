package query

import (
	"context"

	"github.com/google/uuid"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"

	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RoleQueryRepository interface {
	Search(ctx context.Context, q SearchRoleQuery) (SearchRoleQueryResult, *aerrs.AppError)
	Export(ctx context.Context, q ExportRoleQuery) (ExportRoleQueryResult, *aerrs.AppError)
	FindById(ctx context.Context, id uuid.UUID) (*entities.Role, *aerrs.AppError)
	FindByCode(ctx context.Context, code string) (*entities.Role, *aerrs.AppError)
}
