package mapper

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
	"google.golang.org/protobuf/types/known/anypb"
)

// ============================================================
// API
// ============================================================

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

// ============================================================
// DATABASE
// ============================================================

var roleScopeTypeToDBMap = map[enu.RoleScopeType]db.RoleScopeType{
	enu.RoleScopeGobal:  db.RoleScopeType("GLOBAL"),
	enu.RoleScopeTenant: db.RoleScopeType("TENANT"),
	enu.RoleScopeUnit:   db.RoleScopeType("UNIT"),
}

var roleScopeTypeToEntityMap = map[db.RoleScopeType]enu.RoleScopeType{
	db.RoleScopeType("GLOBAL"): enu.RoleScopeGobal,
	db.RoleScopeType("TENANT"): enu.RoleScopeTenant,
	db.RoleScopeType("UNIT"):   enu.RoleScopeUnit,
}

func roleScopeTypeToDB(t enu.RoleScopeType) db.RoleScopeType {
	if v, ok := roleScopeTypeToDBMap[t]; ok {
		return v
	}
	panic(fmt.Sprintf("invalid RoleScopeType: %v", t))
}

func roleScopeTypeFromDB(t db.RoleScopeType) enu.RoleScopeType {
	if v, ok := roleScopeTypeToEntityMap[t]; ok {
		return v
	}
	panic(fmt.Sprintf("invalid RoleScopeType: %v", t))
}

var roleAccessScopeToDBMap = map[enu.RoleAccessScope]db.RoleAccessScope{
	enu.RoleAccessAll: db.RoleAccessScope("ALL"),
	enu.RoleAccessOwn: db.RoleAccessScope("OWN"),
}

var roleAccessScopeToEntityMap = map[db.RoleAccessScope]enu.RoleAccessScope{
	db.RoleAccessScope("ALL"): enu.RoleAccessAll,
	db.RoleAccessScope("OWN"): enu.RoleAccessOwn,
}

func roleAccessScopeToDB(t enu.RoleAccessScope) db.RoleAccessScope {
	if v, ok := roleAccessScopeToDBMap[t]; ok {
		return v
	}
	panic(fmt.Sprintf("invalid RoleAccessScope: %v", t))
}

func roleAccessScopeFromDB(t db.RoleAccessScope) enu.RoleAccessScope {
	if v, ok := roleAccessScopeToEntityMap[t]; ok {
		return v
	}
	panic(fmt.Sprintf("invalid RoleAccessScope: %v", t))
}

func CreateRoleToDBParams(role *entities.Role) db.CreateRoleParams {
	return db.CreateRoleParams{
		ID:              role.ID(),
		ScopeID:         pg.PgUUIDFromUUIDPtr(role.RoleRef().ScopeID()),
		RoleScopeType:   roleScopeTypeToDB(role.RoleScopeType()),
		Code:            role.RoleRef().RoleCode(),
		Name:            role.Name(),
		Description:     pg.TextFromStringPtr(role.Description()),
		RoleAccessScope: roleAccessScopeToDB(role.RoleAccessScope()),
		Level:           pg.Int4FromUint8(role.Level()),
		IsSystem:        role.IsSystem(),
		IsActive:        role.IsActive(),
		IsSuper:         role.IsSuper(),
	}
}

func CreateRoleDBToRoleEntity(row db.CreateRoleRow) entities.Role {
	return entities.NewRoleFromPersistence(
		entities.NewRoleParams{
			ID: row.ID,
			RoleRef: vo.NewRoleRefFromPersistence(
				vo.NewRoleRefParms{
					RoleCode: row.Code,
					ScopeID:  pg.UUIDPtrFromPgUUID(row.ScopeID),
				},
			),
			RoleScopeType:   roleScopeTypeFromDB(row.RoleScopeType),
			Name:            row.Name,
			RoleAccessScope: roleAccessScopeFromDB(row.RoleAccessScope),
			Level:           uint8(row.Level.Int32),
			Description:     pg.StringPtrFromText(row.Description),
			IsSystem:        row.IsSystem,
			IsSuper:         row.IsSuper,
			IsActive:        row.IsActive,
		},
	)
}
