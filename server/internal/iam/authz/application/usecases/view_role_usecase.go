package usecases

import (
	"context"

	coremap "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/mappers"
	coreuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
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
	repo   re.RoleRepository
	engine *coreuc.ViewEngine[RoleView]
}

func NewViewRoleUsecase(repo re.RoleRepository) *ViewRoleUsecase {

	u := &ViewRoleUsecase{
		repo: repo,
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

	results := u.engine.Execute(ctx, batch.Items, coremap.BuildViewResult)

	return dtos.ViewResultDTO{Items: results}
}

func (u *ViewRoleUsecase) handleGet(ctx context.Context, q q.GetRoleQuery) (any, *aerrs.AppError) {
	return nil, nil
}

func (u *ViewRoleUsecase) handleList(ctx context.Context, q q.ListRoleQuery) (any, *aerrs.AppError) {
	return nil, nil
}

func (u *ViewRoleUsecase) handleSearch(ctx context.Context, q q.SearchRoleQuery) (any, *aerrs.AppError) {
	return nil, nil
}

func (u *ViewRoleUsecase) handleExport(ctx context.Context, q q.ExportRoleQuery) (any, *aerrs.AppError) {
	return nil, nil
}
