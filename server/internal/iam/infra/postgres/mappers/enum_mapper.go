package mappers

import (
	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	db "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/sqlc"
)

var roleScopeTypeToDBMap = map[enu.RoleScopeType]db.RoleScopeType{
	enu.RoleScopeGobal:  db.RoleScopeType("GLOBAL"),
	enu.RoleScopeTenant: db.RoleScopeType("TENANT"),
	enu.RoleScopeUnit:   db.RoleScopeType("UNIT"),
}

var roleAccessScopeToDBMap = map[enu.RoleAccessScope]db.RoleAccessScope{
	enu.RoleAccessAll: db.RoleAccessScope("ALL"),
	enu.RoleAccessOwn: db.RoleAccessScope("OWN"),
}

var roleScopeTypeFromDBMap = map[db.RoleScopeType]enu.RoleScopeType{
	db.RoleScopeType("GLOBAL"): enu.RoleScopeGobal,
	db.RoleScopeType("TENANT"): enu.RoleScopeTenant,
	db.RoleScopeType("UNIT"):   enu.RoleScopeUnit,
}

var roleAccessScopeFromDBMap = map[db.RoleAccessScope]enu.RoleAccessScope{
	db.RoleAccessScope("ALL"): enu.RoleAccessAll,
	db.RoleAccessScope("OWN"): enu.RoleAccessOwn,
}

func RoleScopeTypeFromDB(t db.RoleScopeType) enu.RoleScopeType {
	if v, ok := roleScopeTypeFromDBMap[t]; ok {
		return v
	}
	return enu.RoleScopeGobal
}

func RoleAccessScopeFromDB(t db.RoleAccessScope) enu.RoleAccessScope {
	if v, ok := roleAccessScopeFromDBMap[t]; ok {
		return v
	}
	return enu.RoleAccessAll
}

func RoleScopeTypeToDB(t enu.RoleScopeType) db.RoleScopeType {
	if v, ok := roleScopeTypeToDBMap[t]; ok {
		return v
	}
	panic("invalid RoleScopeType from DB: " + string(t))
}

func RoleAccessScopeToDB(t enu.RoleAccessScope) db.RoleAccessScope {
	if v, ok := roleAccessScopeToDBMap[t]; ok {
		return v
	}
	panic("invalid RoleAccessScope from DB: " + string(t))
}
