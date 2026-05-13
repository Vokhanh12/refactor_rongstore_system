package mapper

import (
	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

func RoleToCreateParams(role *en.Role) db.CreateRoleParams {
	return db.CreateRoleParams{
		ID:              role.ID(),
		ScopeID:         pg.PgUUIDFromUUIDPtr(role.RoleKey().ScopeID()),
		RoleScopeType:   RoleScopeTypeToDB(role.ScopeType()),
		Code:            role.RoleKey().RoleCode(),
		Name:            role.Name(),
		Description:     pg.TextFromStringPtr(role.Description()),
		RoleAccessScope: RoleAccessScopeToDB(role.AccessScope()),
		Level:           int32(role.Level()),
		IsSystem:        role.IsSystem(),
		IsActive:        role.IsActive(),
		IsSuper:         role.IsSuper(),
	}
}

func RoleToUpdateParams(role *en.Role) db.UpdateRoleParams {
	return db.UpdateRoleParams{
		ID:              role.ID(),
		ScopeID:         pg.PgUUIDFromUUIDPtr(role.RoleKey().ScopeID()),
		RoleScopeType:   RoleScopeTypeToDB(role.ScopeType()),
		Code:            role.RoleKey().RoleCode(),
		Name:            role.Name(),
		Description:     pg.TextFromStringPtr(role.Description()),
		RoleAccessScope: RoleAccessScopeToDB(role.AccessScope()),
		Level:           int32(role.Level()),
		IsSystem:        role.IsSystem(),
		IsActive:        role.IsActive(),
		IsSuper:         role.IsSuper(),
	}
}

func RoleToExistsByCodeScopeParams(roleScopeType enu.RoleScopeType, RoleKey vo.RoleKey) db.ExistsRoleByCodeScopeParams {
	return db.ExistsRoleByCodeScopeParams{
		Code:          RoleKey.RoleCode(),
		RoleScopeType: RoleScopeTypeToDB(roleScopeType),
		ScopeID:       pg.PgUUIDFromUUIDPtr(RoleKey.ScopeID()),
	}
}

func CreateRoleRowToEntity(row db.CreateRoleRow) en.Role {
	return en.RestoreRole(
		row.ID,
		en.RolePayload{
			RoleKey: vo.RestoreRoleKey(
				row.Code,
				pg.UUIDPtrFromPgUUID(row.ScopeID),
			),
			RoleScopeType:   RoleScopeTypeFromDB(row.RoleScopeType),
			Name:            row.Name,
			RoleAccessScope: RoleAccessScopeFromDB(row.RoleAccessScope),
			Level:           row.Level,
			Description:     pg.StringPtrFromText(row.Description),
			IsSystem:        row.IsSystem,
			IsSuper:         row.IsSuper,
			IsActive:        row.IsActive,
		},
	)
}

func UpdateRoleRowToEntity(row db.UpdateRoleRow) en.Role {
	return en.RestoreRole(
		row.ID,
		en.RolePayload{
			RoleKey: vo.RestoreRoleKey(
				row.Code,
				pg.UUIDPtrFromPgUUID(row.ScopeID),
			),
			RoleScopeType:   RoleScopeTypeFromDB(row.RoleScopeType),
			Name:            row.Name,
			RoleAccessScope: RoleAccessScopeFromDB(row.RoleAccessScope),
			Level:           row.Level,
			Description:     pg.StringPtrFromText(row.Description),
			IsSystem:        row.IsSystem,
			IsSuper:         row.IsSuper,
			IsActive:        row.IsActive,
		},
	)
}
