package usecases

import (
	"context"
	"server/pkg/errors"

	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type RolePermissionMutation struct {
	Create       *c.CreateRolePermissionCommand
	Update       *c.UpdateRolePermissionCommand
	UpdateAvatar *c.UpdateAvatarRolePermissionCommand
	Delete       *c.DeleteRolePermissionCommand
}

type RolePermissionMutationBatch struct {
	Items []u.Operation[RolePermissionMutation]
}

type MutateRolePermissionResult struct {
	Items []common.MutateResultItem
}

type MutateRolePermissionUsecase struct {
	rolePermissionRepository RolePermissionRepository
}

func NewMutateRolePermissionUsecase(rpRepo RolePermissionRepository) *MutateRolePermissionUsecase {
	return &MutateRolePermissionUsecase{
		rolePermissionRepository: rpRepo,
	}
}

func (u *MutateUsecase) Execute(
	ctx context.Context,
	batch RolePermissionMutationBatch,
) *MutateResult {

	ctx, span := otel.Tracer("usecase").
		Start(ctx, "MutateUsecase")
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
			err = u.Update.Handle(ctx, *item.Payload.Update)

		case item.Payload.UpdateAvatar != nil:
			data, err = u.UpdateAvatar.Handle(ctx, *item.Payload.UpdateAvatar)

		case item.Payload.Delete != nil:
			err = u.Delete.Handle(ctx, *item.Payload.Delete)

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

	return &MutateResult{Items: results}
}
