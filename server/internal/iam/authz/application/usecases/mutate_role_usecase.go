package usecases

import (
	"context"
	"server/pkg/errors"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	c "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	re "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/repositories"
	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
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
	Create *c.CreateRoleHandler
	Delete *c.DeleteRoleHandler
	Update *c.UpdateRoleHandler
}

func NewMutateRoleUsecase(roleRepo re.RoleRepository) *MutateRoleUsecase {
	return &MutateRoleUsecase{
		Create: c.NewCreateRoleHandler(roleRepo),
		Delete: c.NewDeleteRoleHandler(roleRepo),
		Update: c.NewUpdateRoleHandler(roleRepo),
	}
}

func (u *MutateRoleUsecase) Execute(
	ctx context.Context,
	batch RoleMutationBatch,
) *common.MutateResult {

	ctx, span := otel.Tracer("usecase").Start(ctx, "MutateUsecase")
	defer span.End()

	results := make([]common.MutateResultItem, 0, len(batch.Items))

	for _, item := range batch.Items {
		var (
			err  *errors.AppError
			data any
		)

		switch {
		case item.Payload.Create != nil:
			data, err = u.Create.Handle(ctx, *item.Payload.Create)

		case item.Payload.Update != nil:
			data, err = u.Update.Handle(ctx, *item.Payload.Update)

		case item.Payload.Delete != nil:
			data, err = u.Delete.Handle(ctx, *item.Payload.Delete)

		default:
			err = errors.New(domain.MUTATE_OPERATION_UNSUPPORTED)
		}

		var code string
		if err != nil {
			code = err.Code

			span.SetAttributes(attribute.Bool("mutate.partial_failure", true))
		}

		results = append(results, common.MutateResultItem{
			OpID:  item.OpID,
			Data:  data,
			Code:  code,
			Error: err,
		})
	}

	return &common.MutateResult{Items: results}
}
