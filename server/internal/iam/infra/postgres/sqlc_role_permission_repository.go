package postgres

import (
	"context"
	"time"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/caches"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	dberrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

var _ caches.RolePermissionCache = (*SqlcRolePermissionRepository)(nil)

type SqlcRolePermissionRepository struct {
	queries *db.Queries
	base    dberrs.BaseRepository
}

// GetPermissions implements RolePermissionCache
func (s *SqlcRolePermissionRepository) GetPermissions(
	ctx context.Context,
	roleCode string,
) ([]string, *apperrors.AppError) {

	perms, err := s.queries.GetRolePermissions(ctx, roleCode)

	if err != nil {
		appErr := s.base.HandleError(err)
		return nil, appErr
	}

	return perms, nil
}

// SetPermissions → bỏ (vì không dùng DB cache)
func (s *SqlcRolePermissionRepository) SetPermissions(
	ctx context.Context,
	roleCode string,
	perms []string,
	ttl time.Duration,
) *apperrors.AppError {
	return nil
}
