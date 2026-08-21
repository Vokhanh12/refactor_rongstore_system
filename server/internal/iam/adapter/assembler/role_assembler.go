package assemblers

import (
	"fmt"

	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	dp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/dispatcher"
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	cif "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/normalize"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	authzuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	"google.golang.org/protobuf/types/known/anypb"
)

func RoleMToUsecase(
	m *authzrs.RoleMutation,
) (dp.Operation, *aerrs.AppError) {

	switch v := m.Action.(type) {

	case *authzrs.RoleMutation_Create:
		if v.Create == nil || v.Create.Data == nil {
			return dp.Operation{}, aerrs.New(core.INVALID_ARGUMENT)
		}

		d := v.Create.Data

		scopeID, err := cif.ParseUUID(d.ScopeId)
		if err != nil {
			return dp.Operation{}, err
		}

		return dp.Operation{
			OpID:   m.OpId,
			Action: authzuc.RoleCreate,
			Payload: &cmd.CreateRoleCommand{
				ScopeID:         scopeID,
				Code:            d.Code,
				RoleScopeType:   d.ScopeType,
				Name:            d.Name,
				Description:     d.Description,
				RoleAccessScope: d.AccessScope,
				Level:           d.Level,
				IsSystem:        d.IsSystem,
				IsActive:        d.IsActive,
				IsSuper:         d.IsSuper,
			},
		}, nil

	case *authzrs.RoleMutation_Update:
		if v.Update == nil || v.Update.Data == nil {
			return dp.Operation{}, aerrs.New(core.INVALID_ARGUMENT)
		}

		d := v.Update.Data

		id, err := cif.ParseUUID(&v.Update.Id)
		if err != nil {
			return dp.Operation{}, err
		}

		scopeID, err := cif.ParseUUID(d.ScopeId)
		if err != nil {
			return dp.Operation{}, err
		}

		return dp.Operation{
			OpID:   m.OpId,
			Action: authzuc.RoleUpdate,
			Payload: &cmd.UpdateRoleCommand{
				ID:              *id,
				ScopeID:         scopeID,
				Code:            d.Code,
				RoleScopeType:   d.ScopeType,
				Name:            d.Name,
				Description:     d.Description,
				RoleAccessScope: d.AccessScope,
				Level:           d.Level,
				IsSystem:        d.IsSystem,
				IsActive:        d.IsActive,
				IsSuper:         d.IsSuper,
			},
		}, nil

	case *authzrs.RoleMutation_Delete:
		if v.Delete == nil {
			return dp.Operation{}, aerrs.New(core.INVALID_ARGUMENT)
		}

		id, err := cif.ParseUUID(&v.Delete.Id)
		if err != nil {
			return dp.Operation{}, err
		}

		return dp.Operation{
			OpID:   m.OpId,
			Action: authzuc.RoleDelete,
			Payload: &cmd.DeleteRoleCommand{
				ID: *id,
			},
		}, nil

	default:
		panic(fmt.Sprintf(
			"unsupported role mutation: %T",
			m.Action,
		))
	}
}

func RoleMToHandler(op dp.Operation) commonv1.MutateResult {
	switch op.Action {
	case authzuc.RoleCreate:
		r := op.Payload.(*cmd.CreateRoleCommandResult)

		return commonv1.MutateResult{
			OpId:       op.OpID,
			ResourceId: "",
			Data:       &anypb.Any{},
			Success:    false,
			ClientErr:  &commonv1.ClientError{},
			ServerErr:  &commonv1.ServerError{},
		}

	case authzuc.RoleUpdate:
		r := op.Payload.(*cmd.UpdateRoleCommandResult)

	case authzuc.RoleDelete:
		r := op.Payload.(*cmd.DeleteRoleCommandResult)
	}
}
