package mappers

import (
	"fmt"

	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	corem "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/mappers"
	coreuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/usecase"
	cmd "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/command"
	authzuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

// ============================================================
// PROTO → COMMAND / QUERY
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

func MapRoleActionRequest(action any) authzuc.RoleMutation {

	switch v := action.(type) {

	case *authzrs.RoleMutation_Create_:
		return authzuc.RoleMutation{
			Create: &cmd.CreateRoleCommand{
				ID: v.Create.Data.Code,
			},
		}

	case *authzrs.RoleMutation_Update_:
		return authzuc.RoleMutation{
			Update: &cmd.UpdateRoleCommand{},
		}

	case *authzrs.RoleMutation_Delete_:
		return authzuc.RoleMutation{
			Delete: &cmd.DeleteRoleCommand{},
		}

	default:
		corem.Must(fmt.Sprintf("unknown action type: %T", action))
		return authzuc.RoleMutation{}
	}
}

func MapRoleActionProto(action any) commonv1.MutateResultItemDTO {

	switch v := action.(type) {

	case *cmd.CreateRoleCommandResult:
		return commonv1.MutateResultItemDTO{
			Data: &authzrs.RoleMutation_Create_{
				Create: &authzrs.RoleMutation_Create{
					Data: &authzrs.RoleMutation_Create_Data{
						Code: v.ID,
					},
				},
			},
		}

	case *cmd.UpdateRoleCommandResult:
		return commonv1.MutateResultItemDTO{
			Data: &authzrs.RoleMutation_Update_{
				Update: &authzrs.RoleMutation_Update{},
			},
		}

	case *cmd.DeleteRoleCommandResult:
		return commonv1.MutateResultItemDTO{
			Data: &authzrs.RoleMutation_Delete_{
				Delete: &authzrs.RoleMutation_Delete{},
			},
		}

	default:
		corem.Must(fmt.Sprintf("unknown action type: %T", action))
		return commonv1.MutateResultItemDTO{}
	}
}
