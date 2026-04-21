package enums

import (
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RoleAccessScope string

const (
	RoleAccessAll RoleAccessScope = "ALL"
	RoleAccessOwn RoleAccessScope = "OWN"
)

func NewRoleAccessScope(v string) (RoleAccessScope, *aerrs.AppError) {
	scope := RoleAccessScope(v)

	if !scope.IsValid() {
		return "", aerrs.New(core.VALIDATION_FAILED,
			aerrs.WithAppendErrorDetail(
				aerrs.NewDetail(core.REASON_VAL_INVALID_FORMAT,
					aerrs.WithField("roleAccessScope"),
					aerrs.WithMessageDetail("invalid role access scope"),
				),
			),
		)
	}

	return scope, nil
}

func (r RoleAccessScope) IsValid() bool {
	switch r {
	case RoleAccessAll, RoleAccessOwn:
		return true
	default:
		return false
	}
}
