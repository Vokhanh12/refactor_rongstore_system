package postgres

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/serialization"
	pr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	authzrepos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var _ authzrepos.AuthorizationQueryRepository = (*AuthorizationQueryRepository)(nil)

type AuthorizationQueryRepository struct {
	dba *pg.DbAdapter
}

func NewAuthorizationQueryRepository(dba *pg.DbAdapter) authzrepos.AuthorizationQueryRepository {
	return &AuthorizationQueryRepository{dba: dba}
}

// ListGrantsByRoleKeys implements [query.AuthorizationQuery].
func (s *AuthorizationQueryRepository) ListGrantsByRoleKeys(ctx context.Context, RoleKeys []valueobjects.RoleKey) ([]pr.AuthorizationGrant, *apperrors.AppError) {

	payload, aerr := serialization.MustMarshal(RoleKeys)
	if aerr != nil {
		return nil, aerr
	}

	rows, err := s.dba.Q.ListAuthorizationGrantsByRoleKeys(ctx, payload)
	if err != nil {
		return nil, s.dba.Translate(err)
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
