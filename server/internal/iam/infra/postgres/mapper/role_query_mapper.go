package mapper

import (
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

func RoleToSearchParams(q q.SearchRoleQuery) db.SearchRolesParams {}
