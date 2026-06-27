package assemblers

import (
	"fmt"

	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/assemblers"
	corem "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/assemblers"
	cif "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/normalize"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	authzuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

// ============================================================
// PROTO → COMMAND (REQUEST)
// ============================================================

func RoleMutateRequestToBatch(req *authzrs.RoleMutateRequest) authzuc.RoleMutationBatch {

	items := assemblers.BuildBatch(
		req.Mutations,
		func(m *authzrs.RoleMutation) (authzuc.RoleMutation, *aerrs.AppError) {
			return DecodeRoleMutation(m.Action)
		},
		func(m *authzrs.RoleMutation) string {
			return m.OpId
		},
	)

	return authzuc.RoleMutationBatch{Items: items}
}

// ============================================================
// ACTION → COMMAND
// ============================================================

func DecodeRoleMutation(action any) (authzuc.RoleMutation, *aerrs.AppError) {

	switch v := action.(type) {

	case *authzrs.RoleMutation_Create:
		scopeID, err := cif.ParseUUID(v.Create.Data.ScopeId)
		if err != nil {
			return authzuc.RoleMutation{}, err
		}
		return authzuc.RoleMutation{
			Create: &cmd.CreateRoleCommand{
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
		id, err := cif.ParseUUID(&v.Update.Id)
		scopeID, err := cif.ParseUUID(v.Update.Data.ScopeId)
		if err != nil {
			return authzuc.RoleMutation{}, err
		}
		return authzuc.RoleMutation{
			Update: &cmd.UpdateRoleCommand{
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
		id, err := cif.ParseUUID(&v.Delete.Id)
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
