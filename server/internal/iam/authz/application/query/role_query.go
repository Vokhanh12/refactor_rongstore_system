package query

import (
	"github.com/google/uuid"
	pr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/querydsl"
	es "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/export"
)

type SearchRoleQuery struct {
	Criteria querydsl.SearchCriteria
}
type SearchRoleQueryResult struct {
	Items []pr.RoleView
	Total int
}

type GetRoleQuery struct {
	ID uuid.UUID
}
type GetRoleQueryResult struct {
	pr.RoleView
}

type ExportRoleQuery struct {
	Criteria querydsl.SearchCriteria
	Format   es.ExportFormat
}

type ExportRoleQueryResult struct {
	FileURL string
}
