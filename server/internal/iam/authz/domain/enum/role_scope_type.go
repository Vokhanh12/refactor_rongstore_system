package enum

import (
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/domain/validator"
)

type RoleScopeType string

const (
	RoleScopeGobal  RoleScopeType = "GLOBAL"
	RoleScopeTenant RoleScopeType = "TENANT"
	RoleScopeUnit   RoleScopeType = "UNIT"
)

var validRoleScopes = map[RoleScopeType]struct{}{
	RoleScopeGobal:  {},
	RoleScopeTenant: {},
	RoleScopeUnit:   {},
}

func NewRoleScopeType(value string) (RoleScopeType, error) {

	v := validator.New().
		Required("RoleScopeType", value)

	scope := RoleScopeType(value)

	v.Enum("roleAccessScope", validator.InEnum(scope, validRoleScopes))

	if err := v.Err(); err != nil {
		return RoleScopeType(""), err
	}

	return scope, nil
}
