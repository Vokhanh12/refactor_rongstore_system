package enum

import (
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/domain/validator"
)

type RoleAccessScope string

const (
	RoleAccessAll RoleAccessScope = "ALL"
	RoleAccessOwn RoleAccessScope = "OWN"
)

var validRoleAccessScopes = map[RoleAccessScope]struct{}{
	RoleAccessAll: {},
	RoleAccessOwn: {},
}

func NewRoleAccessScope(value string) (RoleAccessScope, error) {
	v := validator.New().
		Required("roleAccessScope", value)

	scope := RoleAccessScope(value)

	v.Enum("roleAccessScope", validator.InEnum(scope, validRoleAccessScopes))

	if err := v.Err(); err != nil {
		return RoleAccessScope(""), err
	}

	return scope, nil
}
