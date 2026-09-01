package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/reader"
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

// GetRoleScopesByUserID implements [reader.RoleAssignmentReader].
func (p *PgRoleAsignmentReader) GetRoleScopesByUserID(ctx context.Context, query query.GetRoleScopesQuery) ([]query.GetRoleScopesQueryResult, error) {
	panic("unimplemented")
}
