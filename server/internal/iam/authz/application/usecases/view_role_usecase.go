package usecases

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	q "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/query"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	common "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type RoleView struct {
	Get    *q.GetRoleQuery
	Search *q.SearchRoleQuery
	List   *q.ListRoleQuery
	Export *q.ExportRoleQuery
}

type RoleViewBatch struct {
	Items []core.Operation[RoleView]
}

type ViewRoleUsecase struct {
	repo re.RoleRepository
}

func NewViewRoleUsecase(repo re.RoleRepository) *ViewRoleUsecase {
	return &ViewRoleUsecase{repo: repo}
}

func (u *ViewRoleUsecase) Execute(
	ctx context.Context,
	batch RoleViewBatch,
) *common.ViewResult {

	ctx, span := otel.Tracer("usecase").Start(ctx, "ViewRoleUsecase.Execute")
	defer span.End()

	results := make([]common.ViewResultItem, 0, len(batch.Items))

	for _, item := range batch.Items {
		var (
			err  *aerrs.AppError
			data any
		)

		switch {
		case item.Payload.Get != nil:
			data, err = u.handleGet(ctx, *item.Payload.Get)
		case item.Payload.List != nil:
			data, err = u.handleList(ctx, *item.Payload.List)
		case item.Payload.Search != nil:
			data, err = u.handleSearch(ctx, *item.Payload.Search)
		case item.Payload.Export != nil:
			data, err = u.handleSearch(ctx, *item.Payload.Search)

		default:
			err = aerrs.New(domain.VIEW_OPERATION_UNSUPPORTED)
		}

		var code string
		if err != nil {
			code = err.Code
			span.SetAttributes(attribute.Bool("view.partial_failure", true))
		}

		results = append(results, common.ViewResultItem{
			OpID:  item.OpID,
			Data:  data,
			Code:  code,
			Error: err,
		})
	}

	return &common.ViewResult{Items: results}
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
