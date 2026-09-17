package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/reader"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enum"
	sqlc "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

var _ reader.RoleAssignmentReader = (*PgRoleAsignmentReader)(nil)

type PgRoleAsignmentReader struct {
	queries *sqlc.Queries
	pool    *pgxpool.Pool
}

func NewPgRoleAsignmentReader(q *sqlc.Queries, p *pgxpool.Pool) reader.RoleAssignmentReader {
	return &PgRoleAsignmentReader{queries: q, pool: p}
}

func (p *PgRoleAsignmentReader) ListRoleScopes(
	ctx context.Context,
	userID uuid.UUID,
) ([]query.ListRoleScopesQueryResult, error) {

	rows, err := p.queries.ListRoleScopes(ctx, userID)
	if err != nil {
		return nil, err
	}

	result := make([]query.ListRoleScopesQueryResult, 0, len(rows))

	for _, row := range rows {
		result = append(result, query.ListRoleScopesQueryResult{
			RoleID:    row.RoleID,
			ScopeID:   row.ScopeID,
			ScopeType: enum.RoleScopeType(row.ScopeType),
		})
	}

	return result, nil
}
