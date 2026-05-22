package postgres

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/serialization"
	pr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

var _ q.AuthorizationQuery = (*SqlcAuthorizationQuery)(nil)

type SqlcAuthorizationQuery struct {
	queries *db.Queries
	dberr   dberr.DBError
}

func NewSqlcAuthorizationQuery(queries *db.Queries) q.AuthorizationQuery {
	return &SqlcAuthorizationQuery{queries: queries}
}

// ListGrantsByRoleKeys implements [query.AuthorizationQuery].
func (s *SqlcAuthorizationQuery) ListGrantsByRoleKeys(ctx context.Context, RoleKeys []valueobjects.RoleKey) ([]pr.AuthorizationGrant, *apperrors.AppError) {

	payload, aerr := serialization.MustMarshal(RoleKeys)
	if aerr != nil {
		return nil, aerr
	}

	rows, err := s.queries.ListAuthorizationGrantsByRoleKeys(ctx, payload)
	if err != nil {
		return nil, dberr.TranslateDBError(err, s.dberr)
	}

	results := make([]pr.AuthorizationGrant, 0, len(rows))

	for _, row := range rows {
		results = append(results, pr.AuthorizationGrant{
			RoleKey:        vo.RestoreRoleKey(row.RoleCode, pg.UUIDPtrFromPgUUID(row.RoleScopeID)),
			IsElevated:     row.RoleIsSuper,
			ResourceAction: vo.RestoreResourceAction(row.PermissionResource, row.PermissionAction),
		})
	}

	return results, nil
}
