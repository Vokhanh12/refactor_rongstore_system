package query

import (
	"context"

	"github.com/google/uuid"
	pr "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/projection"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/querydsl"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
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

type RoleQuery interface {
	Search(ctx context.Context, q qs.SearchRoleQuery) (qs.SearchRoleQueryResult, *aerrs.AppError)
	FindById(ctx context.Context, id uuid.UUID) (*entities.Role, *aerrs.AppError)
	FindByCode(ctx context.Context, code string) (*entities.Role, *aerrs.AppError)
	ExistsRoleByCodeScope(ctx context.Context, roleScopeType enu.RoleScopeType, roleKey vo.RoleKey) (bool, *aerrs.AppError)
}
