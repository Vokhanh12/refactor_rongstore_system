package usecases

import (
	"context"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	c "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	common "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type RoleMutation struct {
	Create *c.CreateRoleCommand
	Update *c.UpdateRoleCommand
	Delete *c.DeleteRoleCommand
}

type RoleMutationBatch struct {
	Items []core.Operation[RoleMutation]
}

type MutateRoleUsecase struct {
	repo re.RoleRepository
}

func NewMutateRoleUsecase(repo re.RoleRepository) *MutateRoleUsecase {
	return &MutateRoleUsecase{repo: repo}
}

func (u *MutateRoleUsecase) Execute(ctx context.Context, batch RoleMutationBatch) *common.MutateResult {

	ctx, span := otel.Tracer("usecase").Start(ctx, "MutateRoleUsecase.Execute")
	defer span.End()

	results := make([]common.MutateResultItem, 0, len(batch.Items))

	for _, item := range batch.Items {
		var (
			err  *aerrs.AppError
			data any
		)

		switch {
		case item.Payload.Create != nil:
			data, err = u.handleCreate(ctx, *item.Payload.Create)

		case item.Payload.Update != nil:
			data, err = u.handleUpdate(ctx, *item.Payload.Update)

		case item.Payload.Delete != nil:
			data, err = u.handleDelete(ctx, *item.Payload.Delete)

		default:
			err = aerrs.New(domain.MUTATE_OPERATION_UNSUPPORTED)
		}

		if err != nil {
			span.SetAttributes(attribute.Bool("mutate.partial_failure", true))
		}

		results = append(results, common.MutateResultItem{
			OpID:  item.OpID,
			Data:  data,
			Error: err.ToDTO(),
		})
	}

	return &common.MutateResult{Items: results}
}

func (u *MutateRoleUsecase) handleCreate(ctx context.Context, cmd c.CreateRoleCommand) (*c.CreateRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}

func (u *MutateRoleUsecase) handleUpdate(ctx context.Context, cmd c.UpdateRoleCommand) (*c.UpdateRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}

func (u *MutateRoleUsecase) handleDelete(ctx context.Context, cmd c.DeleteRoleCommand) (*c.DeleteRoleCommandResult, *aerrs.AppError) {
	return nil, nil
}
