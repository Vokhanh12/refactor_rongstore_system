package assemblers

import (
	"fmt"

	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/assemblers"
	dp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/dispatcher"
	cif "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/normalize"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	authzuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

func RoleMutateRequestToBatch(
	req *authzrs.RoleMutateRequest,
) (
	authzuc.RoleMutationBatch,
	*aerrs.AppError,
) {

	items, err := assemblers.BuildBatch(
		req.Mutations,
		decodeRoleMutation,
	)
	if err != nil {
		return authzuc.RoleMutationBatch{}, err
	}

	return authzuc.RoleMutationBatch{
		Items: items,
	}, nil
}

func decodeRoleMutation(
	mutation *authzrs.RoleMutation,
) (
	dp.Operation,
	*aerrs.AppError,
) {

	if mutation == nil {
		return dp.Operation{},
			aerrs.New(core.INVALID_ARGUMENT)
	}

	switch v := mutation.Action.(type) {

	case *authzrs.RoleMutation_Create:
		if v.Create == nil || v.Create.Data == nil {
			return dp.Operation{},
				aerrs.New(core.INVALID_ARGUMENT)
		}

		scopeID, err := cif.ParseUUID(
			v.Create.Data.ScopeId,
		)
		if err != nil {
			return dp.Operation{}, err
		}

		return dp.Operation{
			OpID:   mutation.OpId,
			Action: authzuc.RoleCreate,
			Payload: &cmd.CreateRoleCommand{
				Code:            v.Create.Data.Code,
				ScopeID:         scopeID,
				RoleScopeType:   v.Create.Data.ScopeType,
				Name:            v.Create.Data.Name,
				Description:     v.Create.Data.Description,
				RoleAccessScope: v.Create.Data.AccessScope,
				Level:           v.Create.Data.Level,
				IsSystem:        v.Create.Data.IsSystem,
				IsActive:        v.Create.Data.IsActive,
				IsSuper:         v.Create.Data.IsSuper,
			},
		}, nil

	case *authzrs.RoleMutation_Update:
		if v.Update == nil || v.Update.Data == nil {
			return dp.Operation{},
				aerrs.New(core.INVALID_ARGUMENT)
		}

		id, err := cif.ParseUUID(
			&v.Update.Id,
		)
		if err != nil {
			return dp.Operation{}, err
		}

		scopeID, err := cif.ParseUUID(
			v.Update.Data.ScopeId,
		)
		if err != nil {
			return dp.Operation{}, err
		}

		return dp.Operation{
			OpID:   mutation.OpId,
			Action: authzuc.RoleUpdate,
			Payload: &cmd.UpdateRoleCommand{
				ID:              *id,
				Code:            v.Update.Data.Code,
				ScopeID:         scopeID,
				RoleScopeType:   v.Update.Data.ScopeType,
				Name:            v.Update.Data.Name,
				Description:     v.Update.Data.Description,
				RoleAccessScope: v.Update.Data.AccessScope,
				Level:           v.Update.Data.Level,
				IsSystem:        v.Update.Data.IsSystem,
				IsActive:        v.Update.Data.IsActive,
				IsSuper:         v.Update.Data.IsSuper,
			},
		}, nil

	case *authzrs.RoleMutation_Delete:
		if v.Delete == nil {
			return dp.Operation{},
				aerrs.New(core.INVALID_ARGUMENT)
		}

		id, err := cif.ParseUUID(
			&v.Delete.Id,
		)
		if err != nil {
			return dp.Operation{}, err
		}

		return dp.Operation{
			OpID:   mutation.OpId,
			Action: authzuc.RoleDelete,
			Payload: &cmd.DeleteRoleCommand{
				ID: *id,
			},
		}, nil

	default:
		panic(
			fmt.Sprintf(
				"unsupported role mutation: %T",
				mutation.Action,
			),
		)
	}
}
