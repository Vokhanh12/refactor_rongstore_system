package mapper

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func RolePermissionMutateProtoToCommand(
	req *iamv1.RolePermissionMutateRequest,
) uc.MutationBatch {

	items := make([]ucs.Operation[uc.Mutation], 0, len(req.Mutations))

	for _, m := range req.Mutations {

		op := ucs.Operation[uc.Mutation]{
			OpID: m.OpId,
		}

		switch v := m.Action.(type) {

		case *iamv1.RolePermissionMutation_Create_:
			op.Payload = uc.Mutation{
				Create: &cmds.CreateCommand{
					Lat:   v.Create.Lat,
					Lng:   v.Create.Lng,
					TileX: v.Create.TileX,
					TileY: v.Create.TileY,
				},
			}

		case *iamv1.RolePermissionMutation_Update_:
			op.Payload = uc.Mutation{
				Update: &cmds.UpdateCommand{
					Id:    v.Update.Id,
					Lat:   v.Update.Lat,
					Lng:   v.Update.Lng,
					TileX: v.Update.TileX,
					TileY: v.Update.TileY,
				},
			}

		case *iamv1.RolePermissionMutation_UpdateAvatar_:
			op.Payload = uc.Mutation{
				UpdateAvatar: &cmds.UpdateAvatarCommand{
					OwnerId:     v.UpdateAvatar.OwnerId,
					Data:        v.UpdateAvatar.Data,
					ContentType: v.UpdateAvatar.ContentType,
					Filename:    v.UpdateAvatar.Filename,
				},
			}

		case *iamv1.RolePermissionMutation_Delete_:
			op.Payload = uc.Mutation{
				Delete: &cmds.DeleteCommand{
					Id: v.Delete.Id,
				},
			}

		default:
			op.Payload = uc.Mutation{}
		}

		items = append(items, op)
	}

	return uc.MutationBatch{
		Items: items,
	}
}

func RolePermissionMutateCommandResultToProto(
	result *uc.MutateResult,
) []*commonv1.MutateResult {

	items := make([]*commonv1.MutateResult, 0, len(result.Items))

	for _, item := range result.Items {

		anyData, mapErr := mapMutationDataToAny(item.Data, item.Error)

		items = append(items, &commonv1.MutateResult{
			OpId:    item.OpID,
			Success: item.Success && mapErr == nil,
			Data:    anyData,
			Error:   errors.ToPublicError(mapErr),
		})
	}

	return items
}

func mapMutationDataToAny(
	data any,
	err *errors.AppError,
) (*anypb.Any, *errors.AppError) {

	switch v := data.(type) {

	case *re.CreateRead:
		pb := &iamv1.RolePermissionMutation_Create_Data{
			ResourceId: v.ResourceID,
		}
		return marshalAny(pb)

	case *re.UpdateRead:
		pb := &iamv1.RolePermissionMutation_Update_Data{}
		return marshalAny(pb)

	case *re.UpdateAvatarRead:
		pb := &iamv1.RolePermissionMutation_UpdateAvatar_Data{
			AvatarUrl: v.AvatarUrl,
		}
		return marshalAny(pb)

	case *re.DeleteRead:
		pb := &iamv1.RolePermissionMutation_Delete_Data{}
		return marshalAny(pb)

	case nil:
		return nil, err

	default:
		return nil, errors.New(
			errors.INTERNAL_FALLBACK,
			errors.WithMessage("unsupported mutation data type"),
			errors.WithData(map[string]interface{}{
				"actual_type": fmt.Sprintf("%T", data),
			}),
		)
	}
}

func marshalAny(
	msg proto.Message,
) (*anypb.Any, *errors.AppError) {

	anyMsg, err := anypb.New(msg)
	if err != nil {
		return nil, errors.New(
			errors.INTERNAL_FALLBACK,
			errors.WithCauseDetail(err),
			errors.WithMessage("failed to marshal mutation data to Any"),
		)
	}

	return anyMsg, nil
}
