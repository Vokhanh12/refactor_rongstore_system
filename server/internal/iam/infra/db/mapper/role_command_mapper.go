package mapper

import (
	"github.com/google/uuid"

	en "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"

	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

func CreateRoleRowToEntity(
	row db.CreateRoleRow,
) en.Role {
	return restoreRoleEntity(
		row.ID,
		row.ScopeID,
		row.RoleScopeType,
		row.Code,
		row.Name,
		row.Description,
		row.RoleAccessScope,
		row.Level,
		row.IsSystem,
		row.IsActive,
		row.IsSuper,
	)
}

func UpdateRoleRowToEntity(
	row db.UpdateRoleRow,
) en.Role {
	return restoreRoleEntity(
		row.ID,
		row.ScopeID,
		row.RoleScopeType,
		row.Code,
		row.Name,
		row.Description,
		row.RoleAccessScope,
		row.Level,
		row.IsSystem,
		row.IsActive,
		row.IsSuper,
	)
}

func restoreRoleEntity(
	id uuid.UUID,
	scopeID *uuid.UUID,
	scopeType db.RoleScopeType,
	code string,
	name string,
	description *string,
	accessScope db.RoleAccessScope,
	level int32,
	isSystem bool,
	isActive bool,
	isSuper bool,
) en.Role {

	return en.RestoreRole(
		id,
		en.RolePayload{
			RoleKey: vo.RestoreRoleKey(
				code,
				scopeID,
			),

			RoleScopeType: RoleScopeTypeFromDB(
				scopeType,
			),

			Name: name,

			RoleAccessScope: RoleAccessScopeFromDB(
				accessScope,
			),

			Level: level,

			Description: description,

			IsSystem: isSystem,
			IsSuper:  isSuper,
			IsActive: isActive,
		},
	)
}
