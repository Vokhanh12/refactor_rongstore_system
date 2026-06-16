package usecases

import (
	"context"

	coreuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/usecase"
	qs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	repos "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type RoleView struct {
	Get    *qs.GetRoleQuery
	Search *qs.SearchRoleQuery
	Export *qs.ExportRoleQuery
}

type RoleViewBatch struct {
	Items []coreuc.Operation[RoleView]
}

type ViewRoleUsecase struct {
	query  repos.RoleQueryRepository
	engine *coreuc.ViewEngine[RoleView]
}

func NewViewRoleUsecase(q repos.RoleQueryRepository) *ViewRoleUsecase {

	u := &ViewRoleUsecase{
		query: q,
	}

	handlers := []coreuc.Handler[RoleView]{
		{
			Cond: func(p RoleView) bool { return p.Export != nil },
			Exec: func(ctx context.Context, p RoleView) (any, *aerrs.AppError) {
				return u.handleExport(ctx, *p.Export)
			},
		},
		{
			Cond: func(p RoleView) bool { return p.Get != nil },
			Exec: func(ctx context.Context, p RoleView) (any, *aerrs.AppError) {
				return u.handleGet(ctx, *p.Get)
			},
		},
		{
			Cond: func(p RoleView) bool { return p.Search != nil },
			Exec: func(ctx context.Context, p RoleView) (any, *aerrs.AppError) {
				return u.handleSearch(ctx, *p.Search)
			},
		},
	}

	u.engine = coreuc.NewViewEngine(handlers)

	return u
}

func (u *ViewRoleUsecase) Execute(ctx context.Context, batch RoleViewBatch) dtos.ViewResultDTO {

	results := u.engine.Execute(ctx, batch.Items)

	return dtos.ViewResultDTO{Items: results}
}

func (u *ViewRoleUsecase) handleGet(ctx context.Context, q qs.GetRoleQuery) (qs.GetRoleQueryResult, *aerrs.AppError) {
	return u.query.Get(ctx, q)
}

func (u *ViewRoleUsecase) handleSearch(ctx context.Context, q qs.SearchRoleQuery) (qs.SearchRoleQueryResult, *aerrs.AppError) {
	return u.query.Search(ctx, q)
}

func (u *ViewRoleUsecase) handleExport(ctx context.Context, q qs.ExportRoleQuery) (qs.ExportRoleQueryResult, *aerrs.AppError) {
	return u.query.Export(ctx, q)
}
