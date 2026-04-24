package enums

import (
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/validator"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
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

func NewRoleAccessScope(value string) (RoleAccessScope, *aerrs.AppError) {
	v := validator.New().
		Required("roleAccessScope", value)

	scope := RoleAccessScope(value)

	v.Enum("roleAccessScope", validator.InEnum(scope, validRoleAccessScopes))

	if err := v.Err(); err != nil {
		return "", err
	}

	return scope, nil
}
