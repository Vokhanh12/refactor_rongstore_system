package mappers

import (
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

func RoleToCreateParams(role *entities.Role) db.CreateRoleParams {
	return db.CreateRoleParams{
		ID:              role.ID(),
		ScopeID:         pg.PgUUIDFromUUIDPtr(role.RoleRef().ScopeID()),
		RoleScopeType:   RoleScopeTypeToDB(role.RoleScopeType()),
		Code:            role.RoleRef().RoleCode(),
		Name:            role.Name(),
		Description:     pg.TextFromStringPtr(role.Description()),
		RoleAccessScope: RoleAccessScopeToDB(role.RoleAccessScope()),
		Level:           int32(role.Level()),
		IsSystem:        role.IsSystem(),
		IsActive:        role.IsActive(),
		IsSuper:         role.IsSuper(),
	}
}

func RoleToUpdateParams(role *entities.Role) db.UpdateRoleParams {
	return db.UpdateRoleParams{
		ID:              role.ID(),
		ScopeID:         pg.PgUUIDFromUUIDPtr(role.RoleRef().ScopeID()),
		RoleScopeType:   RoleScopeTypeToDB(role.RoleScopeType()),
		Code:            role.RoleRef().RoleCode(),
		Name:            role.Name(),
		Description:     pg.TextFromStringPtr(role.Description()),
		RoleAccessScope: RoleAccessScopeToDB(role.RoleAccessScope()),
		Level:           int32(role.Level()),
		IsSystem:        role.IsSystem(),
		IsActive:        role.IsActive(),
		IsSuper:         role.IsSuper(),
	}
}

func CreateRoleRowToEntity(row db.CreateRoleRow) entities.Role {
	return entities.NewRoleFromPersistence(
		entities.NewRoleParams{
			ID: row.ID,
			RoleRef: vo.RestoreRoleRef(
				row.Code,
				pg.UUIDPtrFromPgUUID(row.ScopeID),
			),
			RoleScopeType:   RoleScopeTypeFromDB(row.RoleScopeType),
			Name:            row.Name,
			RoleAccessScope: RoleAccessScopeFromDB(row.RoleAccessScope),
			Level:           pg.Uint8FromInt32(row.Level),
			Description:     pg.StringPtrFromText(row.Description),
			IsSystem:        row.IsSystem,
			IsSuper:         row.IsSuper,
			IsActive:        row.IsActive,
		},
	)
}

func UpdateRoleRowToEntity(row db.UpdateRoleRow) entities.Role {
	return entities.NewRoleFromPersistence(
		entities.NewRoleParams{
			ID: row.ID,
			RoleRef: vo.RestoreRoleRef(
				row.Code,
				pg.UUIDPtrFromPgUUID(row.ScopeID),
			),
			RoleScopeType:   RoleScopeTypeFromDB(row.RoleScopeType),
			Name:            row.Name,
			RoleAccessScope: RoleAccessScopeFromDB(row.RoleAccessScope),
			Level:           pg.Uint8FromInt32(row.Level),
			Description:     pg.StringPtrFromText(row.Description),
			IsSystem:        row.IsSystem,
			IsSuper:         row.IsSuper,
			IsActive:        row.IsActive,
		},
	)
}
