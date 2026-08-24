package scanrows

import (
	"github.com/jackc/pgx/v5"
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
)

func ScanRoleView(rows pgx.Rows) (q.RoleView, error) {

	var item q.RoleView

	err := rows.Scan(
		&item.ID,
		&item.Code,
		&item.Name,
		&item.ScopeType,
		&item.Level,
		&item.Description,
		&item.IsSystem,
		&item.IsSuper,
		&item.IsActive,
	)

	return item, err
}
