package postgres

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/serialization"
	pr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ q.AuthorizationQuery = (*SqlcAuthorizationQuery)(nil)

type SqlcAuthorizationQuery struct {
	dba *pg.DbAdapter
}

func NewSqlcAuthorizationQuery(dba *pg.DbAdapter) q.AuthorizationQuery {
	return &SqlcAuthorizationQuery{dba: dba}
}

// ListGrantsByRoleKeys implements [query.AuthorizationQuery].
func (s *SqlcAuthorizationQuery) ListGrantsByRoleKeys(ctx context.Context, RoleKeys []valueobjects.RoleKey) ([]pr.AuthorizationGrant, *apperrors.AppError) {

	payload, aerr := serialization.MustMarshal(RoleKeys)
	if aerr != nil {
		return nil, aerr
	}

	rows, err := s.dba.Q.ListAuthorizationGrantsByRoleKeys(ctx, payload)
	if err != nil {
		return nil, s.dba.Wrap(err)
	}

	results := make([]pr.AuthorizationGrant, 0, len(rows))

	for _, row := range rows {
		results = append(results, pr.AuthorizationGrant{
			RoleKey:        vo.RestoreRoleKey(row.RoleCode, row.RoleScopeID),
			IsElevated:     row.RoleIsSuper,
			ResourceAction: vo.RestoreResourceAction(row.PermissionResource, row.PermissionAction),
		})
	}

	return results, nil
}
