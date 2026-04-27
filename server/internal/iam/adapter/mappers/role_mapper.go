package mappers

import (
	"fmt"

	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	corem "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/mappers"
	coreuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	authzuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
)

// ============================================================
// PROTO → COMMAND (REQUEST)
// ============================================================

func RoleMutateRequestToBatch(req *authzrs.RoleMutateRequest) authzuc.RoleMutationBatch {

	items := make([]coreuc.Operation[authzuc.RoleMutation], 0, len(req.Mutations))

	for _, m := range req.Mutations {
		items = append(items, coreuc.Operation[authzuc.RoleMutation]{
			OpID:    m.OpId,
			Payload: MapRoleActionRequest(m.Action),
		})
	}

	return authzuc.RoleMutationBatch{Items: items}
}

// ============================================================
// ACTION → COMMAND
// ============================================================

func MapRoleActionRequest(action any) authzuc.RoleMutation {

	switch v := action.(type) {

	case *authzrs.RoleMutation_Create_:
		return authzuc.RoleMutation{
			Create: &cmd.CreateRoleCommand{
				Code:            v.Create.Data.Code,
				ScopeID:         v.Create.Data.ScopeId,
				RoleScopeType:   v.Create.Data.ScopeType,
				Name:            v.Create.Data.Name,
				Description:     v.Create.Data.Description,
				RoleAccessScope: v.Create.Data.AccessScope,
				Level:           v.Create.Data.Level,
				IsSystem:        v.Create.Data.IsSystem, // FIX BUG
				IsActive:        v.Create.Data.IsActive,
				IsSuper:         v.Create.Data.IsSuper,
			},
		}

	case *authzrs.RoleMutation_Update_:
		return authzuc.RoleMutation{
			Update: &cmd.UpdateRoleCommand{
				ID:              v.Update.Id,
				Code:            v.Update.Data.Code,
				ScopeID:         v.Update.Data.ScopeId,
				RoleScopeType:   v.Update.Data.ScopeType,
				Name:            v.Update.Data.Name,
				Description:     v.Update.Data.Description,
				RoleAccessScope: v.Update.Data.AccessScope,
				Level:           v.Update.Data.Level,
				IsSystem:        v.Update.Data.IsSystem, // FIX BUG
				IsActive:        v.Update.Data.IsActive,
				IsSuper:         v.Update.Data.IsSuper,
			},
		}

	case *authzrs.RoleMutation_Delete_:
		return authzuc.RoleMutation{
			Delete: &cmd.DeleteRoleCommand{
				ID: v.Delete.Id,
			},
		}

	default:
		corem.Must(fmt.Sprintf("unknown action type: %T", action))
		return authzuc.RoleMutation{}
	}
}

// ============================================================
// COMMAND RESULT → PROTO RESPONSE (DATA PART)
// ============================================================

func MapRoleActionResponse(data any) authzrs.RoleMutation {

	switch v := data.(type) {

	case *cmd.CreateRoleCommandResult:
		return &authzrs.MutateResultItem_Create{
			Create: &authzrs.CreateRoleResult{
				Id: v.Result.Id,
			},
		}

	case *cmd.UpdateRoleCommandResult:
		return &authzrs.MutateResultItem_Update{
			Update: &authzrs.UpdateRoleResult{},
		}

	case *cmd.DeleteRoleCommandResult:
		return &authzrs.MutateResultItem_Delete{
			Delete: &authzrs.DeleteRoleResult{},
		}

	default:
		corem.Must(fmt.Sprintf("unknown result type: %T", data))
		return nil
	}
}
