package mappers

import (
	authzrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/authz/v1/resources"
	authzuc "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/application/usecases"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/entities"
	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	pg "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/postgres"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

// ============================================================
// INTERFACE
// ============================================================

type AuthzMapper interface {
	// proto → command
	ToRoleMutationBatch(req *authzrs.RoleMutateRequest) authzuc.RoleMutationBatch
	ToPermissionMutationBatch(req *authzrs.PermissionMutateRequest) authzuc.PermissionMutationBatch
	ToRolePermissionMutationBatch(req *authzrs.RolePermissionMutateRequest) authzuc.RolePermissionMutationBatch

	// proto → query
	ToRoleViewBatch(req *authzrs.RoleViewRequest) authzuc.RoleViewBatch
	ToPermissionViewBatch(req *authzrs.PermissionViewRequest) authzuc.PermissionViewBatch
	ToRolePermissionViewBatch(req *authzrs.RolePermissionViewRequest) authzuc.RolePermissionViewBatch

	// entity → db
	ToCreateRoleParams(role *entities.Role) db.CreateRoleParams

	// db → entity
	ToRoleEntity(row db.CreateRoleRow) entities.Role

	// enum mapping
	ToRoleScopeType(t db.RoleScopeType) enu.RoleScopeType
	ToRoleAccessScopeDB(t enu.RoleAccessScope) db.RoleAccessScope
}

// compile check
var _ AuthzMapper = (*DefaultAuthzMapper)(nil)

// ============================================================
// IMPLEMENTATION
// ============================================================

type DefaultAuthzMapper struct{}

// ============================================================
// ENUM MAPPING
// ============================================================

var roleScopeTypeToDBMap = map[enu.RoleScopeType]db.RoleScopeType{
	enu.RoleScopeGobal:  db.RoleScopeType("GLOBAL"),
	enu.RoleScopeTenant: db.RoleScopeType("TENANT"),
	enu.RoleScopeUnit:   db.RoleScopeType("UNIT"),
}

var roleScopeTypeFromDBMap = map[db.RoleScopeType]enu.RoleScopeType{
	db.RoleScopeType("GLOBAL"): enu.RoleScopeGobal,
	db.RoleScopeType("TENANT"): enu.RoleScopeTenant,
	db.RoleScopeType("UNIT"):   enu.RoleScopeUnit,
}

func (m *DefaultAuthzMapper) ToRoleScopeType(t db.RoleScopeType) enu.RoleScopeType {
	if v, ok := roleScopeTypeFromDBMap[t]; ok {
		return v
	}
	return enu.RoleScopeGobal
}

func (m *DefaultAuthzMapper) ToRoleAccessScopeDB(t enu.RoleAccessScope) db.RoleAccessScope {
	switch t {
	case enu.RoleAccessAll:
		return db.RoleAccessScope("ALL")
	case enu.RoleAccessOwn:
		return db.RoleAccessScope("OWN")
	default:
		return db.RoleAccessScope("ALL")
	}
}

// ============================================================
// ENTITY → DB
// ============================================================

func (m *DefaultAuthzMapper) ToCreateRoleParams(role *entities.Role) db.CreateRoleParams {
	return db.CreateRoleParams{
		ID:              role.ID(),
		ScopeID:         pg.PgUUIDFromUUIDPtr(role.RoleRef().ScopeID()),
		RoleScopeType:   roleScopeTypeToDBMap[role.RoleScopeType()],
		Code:            role.RoleRef().RoleCode(),
		Name:            role.Name(),
		Description:     pg.TextFromStringPtr(role.Description()),
		RoleAccessScope: m.ToRoleAccessScopeDB(role.RoleAccessScope()),
		Level:           pg.Int4FromUint8(role.Level()),
		IsSystem:        role.IsSystem(),
		IsActive:        role.IsActive(),
		IsSuper:         role.IsSuper(),
	}
}

// ============================================================
// DB → ENTITY
// ============================================================

func (m *DefaultAuthzMapper) ToRoleEntity(row db.CreateRoleRow) entities.Role {
	return entities.NewRoleFromPersistence(
		entities.NewRoleParams{
			ID: row.ID,
			RoleRef: vo.NewRoleRefFromPersistence(vo.NewRoleRefParms{
				RoleCode: row.Code,
				ScopeID:  pg.UUIDPtrFromPgUUID(row.ScopeID),
			}),
			RoleScopeType:   m.ToRoleScopeType(row.RoleScopeType),
			Name:            row.Name,
			RoleAccessScope: enu.RoleAccessScope(row.RoleAccessScope),
			Level:           pg.Uint8FromInt4(row.Level),
			Description:     pg.StringPtrFromText(row.Description),
			IsSystem:        row.IsSystem,
			IsSuper:         row.IsSuper,
			IsActive:        row.IsActive,
		},
	)
}

// ============================================================
// PROTO → COMMAND / QUERY (SKELETON)
// ============================================================

// NOTE: bạn implement dần theo business, giữ skeleton này

func (m *DefaultAuthzMapper) ToRoleMutationBatch(
	req *authzrs.RoleMutateRequest,
) authzuc.RoleMutationBatch {
	// TODO: implement giống pattern bạn đang làm
	return authzuc.RoleMutationBatch{}
}

func (m *DefaultAuthzMapper) ToPermissionMutationBatch(
	req *authzrs.PermissionMutateRequest,
) authzuc.PermissionMutationBatch {
	return authzuc.PermissionMutationBatch{}
}

func (m *DefaultAuthzMapper) ToRolePermissionMutationBatch(
	req *authzrs.RolePermissionMutateRequest,
) authzuc.RolePermissionMutationBatch {
	return authzuc.RolePermissionMutationBatch{}
}

func (m *DefaultAuthzMapper) ToRoleViewBatch(
	req *authzrs.RoleViewRequest,
) authzuc.RoleViewBatch {
	return authzuc.RoleViewBatch{}
}

func (m *DefaultAuthzMapper) ToPermissionViewBatch(
	req *authzrs.PermissionViewRequest,
) authzuc.PermissionViewBatch {
	return authzuc.PermissionViewBatch{}
}

func (m *DefaultAuthzMapper) ToRolePermissionViewBatch(
	req *authzrs.RolePermissionViewRequest,
) authzuc.RolePermissionViewBatch {
	return authzuc.RolePermissionViewBatch{}
}
