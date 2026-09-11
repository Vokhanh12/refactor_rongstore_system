package mapper

import (
	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"

	dp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/dispatcher"
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	cif "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/normalize"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	authzuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func ToMutateRoleCommand(
	m *authzrs.RoleMutateRequest_Mutation,
) (dp.Operation, error) {
	switch action := m.Action.(type) {
	case *authzrs.RoleMutateRequest_Mutation_Create_:
		return mapCreateRole(m, action.Create)

	case *authzrs.RoleMutateRequest_Mutation_Update_:
		return mapUpdateRole(m, action.Update)

	case *authzrs.RoleMutateRequest_Mutation_Delete_:
		return mapDeleteRole(m, action.Delete)

	default:
		return dp.Operation{}, aerrs.New(core.INVALID_ARGUMENT)
	}
}

func mapCreateRole(
	m *authzrs.RoleMutateRequest_Mutation,
	req *authzrs.RoleMutateRequest_Mutation_Create,
) (dp.Operation, error) {
	if req == nil || req.Data == nil {
		return dp.Operation{}, aerrs.New(core.INVALID_ARGUMENT)
	}

	data := req.Data

	scopeID, err := cif.ParseUUID(data.ScopeId)
	if err != nil {
		return dp.Operation{}, err
	}

	return dp.Operation{
		OpID:   m.OpId,
		Action: authzuc.RoleCreate,
		Payload: &cmd.CreateRoleCommand{
			ScopeID:         scopeID,
			Code:            data.Code,
			RoleScopeType:   data.ScopeType,
			Name:            data.Name,
			Description:     data.Description,
			RoleAccessScope: data.AccessScope,
			Level:           data.Level,
			IsSystem:        data.IsSystem,
			IsActive:        data.IsActive,
			IsSuper:         data.IsSuper,
		},
	}, nil
}

func mapUpdateRole(
	m *authzrs.RoleMutateRequest_Mutation,
	req *authzrs.RoleMutateRequest_Mutation_Update,
) (dp.Operation, error) {
	if req == nil || req.Data == nil {
		return dp.Operation{}, aerrs.New(core.INVALID_ARGUMENT)
	}

	id, err := cif.ParseUUID(req.Id)
	if err != nil {
		return dp.Operation{}, err
	}

	scopeID, err := cif.ParseUUID(req.Data.ScopeId)
	if err != nil {
		return dp.Operation{}, err
	}

	data := req.Data

	return dp.Operation{
		OpID:   m.OpId,
		Action: authzuc.RoleUpdate,
		Payload: &cmd.UpdateRoleCommand{
			ID:              *id,
			ScopeID:         scopeID,
			Code:            data.Code,
			RoleScopeType:   data.ScopeType,
			Name:            data.Name,
			Description:     data.Description,
			RoleAccessScope: data.AccessScope,
			Level:           data.Level,
			IsSystem:        data.IsSystem,
			IsActive:        data.IsActive,
			IsSuper:         data.IsSuper,
		},
	}, nil
}

func mapDeleteRole(
	m *authzrs.RoleMutateRequest_Mutation,
	req *authzrs.RoleMutateRequest_Mutation_Delete,
) (dp.Operation, error) {
	if req == nil {
		return dp.Operation{}, aerrs.New(core.INVALID_ARGUMENT)
	}

	id, err := cif.ParseUUID(req.Id)
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
}

// FromMutateRoleResult maps an application command result
// into the generic transport mutation result.
func FromMutateRoleResult(
	op dp.Operation,
	result any,
) (commonv1.MutateResult, error) {
	var data proto.Message

	switch op.Action {
	case authzuc.RoleCreate:
		r, ok := result.(*cmd.CreateRoleCommandResult)
		if !ok {
			return commonv1.MutateResult{}, aerrs.New(core.INTERNAL)
		}

		data = &authzrs.RoleMutationResult_CreateResult{
			Id: r.Role.Id,
		}

	case authzuc.RoleUpdate:
		r, ok := result.(*cmd.UpdateRoleCommandResult)
		if !ok {
			return commonv1.MutateResult{}, aerrs.New(core.INTERNAL)
		}

		data = &authzrs.RoleMutationResult_UpdateResult{
			Id: r.Role.Id,
		}

	case authzuc.RoleDelete:
		r, ok := result.(*cmd.DeleteRoleCommandResult)
		if !ok {
			return commonv1.MutateResult{}, aerrs.New(core.INTERNAL)
		}

		data = &authzrs.RoleMutationResult_DeleteResult{
			Id: r.Role.Id,
		}

	default:
		return commonv1.MutateResult{}, aerrs.New(core.INVALID_ARGUMENT)
	}

	anyData, err := anypb.New(data)
	if err != nil {
		return commonv1.MutateResult{}, err
	}

	return commonv1.MutateResult{
		Metadata: &commonv1.MetadataReponse{
			OpId: op.OpID,
		},
		Data: anyData,
	}, nil
}
