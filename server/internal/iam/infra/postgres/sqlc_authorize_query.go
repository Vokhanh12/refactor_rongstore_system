package postgres

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/serialization"
	pr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	dberr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/errors"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ q.AuthorizeQuery = (*SqlcAuthorizeQuery)(nil)

type SqlcAuthorizeQuery struct {
	queries *db.Queries
	dberr   dberr.DBError
}

func NewSqlcAuthorizeQuery(queries *db.Queries) q.AuthorizeQuery {
	return &SqlcAuthorizeQuery{queries: queries}
}

// ListGrantsByRoleRefs implements [query.AuthorizeQuery].
func (s *SqlcAuthorizeQuery) ListGrantsByRoleRefs(ctx context.Context, roleRefs []valueobjects.RoleRef) ([]pr.AuthorizationGrant, *apperrors.AppError) {

	payload, aerr := serialization.MustMarshal(roleRefs)
	if aerr != nil {
		return nil, aerr
	}

	rows, err := s.queries.ListAuthorizationGrantsByRoleRefs(ctx, payload)
	if err != nil {
		return nil, dberr.TranslateDBError(err, s.dberr)
	}

	results := make([]pr.AuthorizationGrant, 0, len(rows))

	for _, row := range rows {
		results = append(results, pr.AuthorizationGrant{
			RoleRef:    vo.RestoreRoleRef(row.RoleCode, pg.UUIDPtrFromPgUUID(row.RoleScopeID)),
			IsElevated: row.RoleIsSuper,
			Resource:   row.PermissionResource,
			Action:     row.PermissionAction,
		})
	}

	return results, nil
}
