package mappers

import (
	"fmt"

	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	corem "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/mappers"
	cif "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra"
	coreuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	authzuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	"google.golang.org/protobuf/types/known/anypb"
)

// ============================================================
// PROTO → COMMAND (REQUEST)
// ============================================================

func RoleMutateRequestToBatch(req *authzrs.RoleMutateRequest) authzuc.RoleMutationBatch {

	items := make([]coreuc.Operation[authzuc.RoleMutation], 0, len(req.Mutations))

	for _, m := range req.Mutations {
		payload, err := DecodeRoleMutation(m.Action)

		if err != nil {
			items = append(items, coreuc.Operation[authzuc.RoleMutation]{
				OpID:    m.OpId,
				Payload: payload,
				Success: false,
			})
		}

		items = append(items, coreuc.Operation[authzuc.RoleMutation]{
			OpID:    m.OpId,
			Payload: payload,
			Success: true,
		})
	}

	return authzuc.RoleMutationBatch{Items: items}
}

// ============================================================
// ACTION → COMMAND
// ============================================================

func DecodeRoleMutation(action any) (authzuc.RoleMutation, *aerrs.AppError) {

	switch v := action.(type) {

	case *authzrs.RoleMutation_Create:
		scopeID, err := cif.UUIDParse(v.Create.Data.ScopeId)
		if err != nil {
			return authzuc.RoleMutation{}, err
		}
		return authzuc.RoleMutation{
			Create: &cmd.CreateRoleCommand{
				Code:            v.Create.Data.Code,
				ScopeID:         *scopeID,
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
		id, err := cif.UUIDParse(v.Update.Id)
		if err != nil {
			return authzuc.RoleMutation{}, err
		}
		return authzuc.RoleMutation{
			Update: &cmd.UpdateRoleCommand{
				ID:              *id,
				Code:            v.Update.Data.Code,
				ScopeID:         v.Update.Data.ScopeId,
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
		id, err := cif.UUIDParse(v.Delete.Id)
		if err != nil {
			return authzuc.RoleMutation{}, err
		}
		return authzuc.RoleMutation{
			Delete: &cmd.DeleteRoleCommand{
				ID: *id,
			},
		}, nil

	default:
		corem.Must(fmt.Sprintf("unknown action type: %T", action))
		return authzuc.RoleMutation{}, nil
	}
}

// ============================================================
// COMMAND RESULT → PROTO RESPONSE (DATA PART)
// ============================================================

func EncoreRoleMutation(data any) *anypb.Any {

	switch v := data.(type) {

	case *cmd.CreateRoleCommandResult:
		pb := &authzrs.CreateResult{
			RoleResult: &authzrs.RoleResult{
				Id:          v.Result.Id,
				Name:        v.Result.Name,
				Description: v.Result.Description,
				CreateAt:    corem.ToProtoTime(v.Result.CreatedAt),
				UpdateAt:    corem.ToProtoTime(v.Result.UpdatedAt),
			},
		}
		return corem.MustMarshalAny(pb)

	case *cmd.UpdateRoleCommandResult:
		pb := &authzrs.UpdateResult{
			RoleResult: &authzrs.RoleResult{
				Id:          v.Result.Id,
				Name:        v.Result.Name,
				Description: v.Result.Description,
				CreateAt:    corem.ToProtoTime(v.Result.UpdatedAt),
				UpdateAt:    corem.ToProtoTime(v.Result.UpdatedAt),
			},
		}
		return corem.MustMarshalAny(pb)

	case *cmd.DeleteRoleCommandResult:
		pb := &authzrs.DeleteResult{}
		return corem.MustMarshalAny(pb)

	default:
		corem.Must(fmt.Sprintf("unknown result type: %T", data))
		return nil
	}
}
