package usecases

import (
	"context"

	coreuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/usecase"
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

type RoleView struct {
	Get    *q.GetRoleQuery
	Search *q.SearchRoleQuery
	List   *q.ListRoleQuery
	Export *q.ExportRoleQuery
}

type RoleViewBatch struct {
	Items []coreuc.Operation[RoleView]
}

type ViewRoleUsecase struct {
	query  q.RoleQuery
	engine *coreuc.ViewEngine[RoleView]
}

func NewViewRoleUsecase(q q.RoleQuery) *ViewRoleUsecase {

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
			Cond: func(p RoleView) bool { return p.List != nil },
			Exec: func(ctx context.Context, p RoleView) (any, *aerrs.AppError) {
				return u.handleList(ctx, *p.List)
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

func (u *ViewRoleUsecase) handleGet(ctx context.Context, q q.GetRoleQuery) (q.GetRoleQueryResult, *aerrs.AppError) {
	return u.query.Get(ctx, q)
}

func (u *ViewRoleUsecase) handleList(ctx context.Context, q q.ListRoleQuery) (q.ListRoleQueryResult, *aerrs.AppError) {
	return u.query.List(ctx, q)
}

func (u *ViewRoleUsecase) handleSearch(ctx context.Context, q q.SearchRoleQuery) (q.SearchRoleQueryResult, *aerrs.AppError) {
	return u.query.Search(ctx, q)
}

func (u *ViewRoleUsecase) handleExport(ctx context.Context, q q.ExportRoleQuery) (q.ExportRoleQueryResult, *aerrs.AppError) {
	return u.query.Export(ctx, q)
}
