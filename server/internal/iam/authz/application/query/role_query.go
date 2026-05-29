package query

import (
	pr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/querydsl"
)

type SearchRoleQuery struct {
	Criteria querydsl.SearchCriteria
}
type SearchRoleQueryResult struct {
	Items []pr.RoleView
	Total int
}

type ListRoleQuery struct{}
type ListRoleQueryResult struct{}

type GetRoleQuery struct{}
type GetRoleQueryResult struct{}

type ExportRoleQuery struct{}
type ExportRoleQueryResult struct{}

type CountRoleQuery struct{}
type CountRoleQueryResult struct{}

type ExistsRoleQuery struct{}
type ExistsRoleQueryResult struct{}

type ListRoleByRefQuery struct{}
type ListRoleByRefQueryResult struct{}
