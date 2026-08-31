package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/serialization"
	pr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	authzrepos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	sqlc "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

var _ authzrepos.AuthorizationQueryRepository = (*AuthorizationQueryRepository)(nil)

type AuthorizationQueryRepository struct {
	queries *sqlc.Queries
	pool    *pgxpool.Pool
}

func NewAuthorizationQueryRepository(q *sqlc.Queries, p *pgxpool.Pool) authzrepos.AuthorizationQueryRepository {
	return &AuthorizationQueryRepository{
		queries: q,
		pool:    p,
	}
}

// ListGrantsByRoleKeys implements [query.AuthorizationQuery].
func (s *AuthorizationQueryRepository) ListGrantsByRoleKeys(ctx context.Context, RoleKeys []vo.RoleKey) ([]pr.AuthorizationGrant, error) {

	payload, aerr := serialization.MustMarshal(RoleKeys)
	if aerr != nil {
		return nil, aerr
	}

	rows, err := s.queries.ListAuthorizationGrantsByRoleKeys(ctx, payload)
	if err != nil {
		return nil, err
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
