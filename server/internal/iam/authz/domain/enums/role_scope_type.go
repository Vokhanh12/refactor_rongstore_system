package enums

import (
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type RoleScopeType string

const (
	RoleScopeGobal  RoleScopeType = "GOBAL"
	RoleScopeTenant RoleScopeType = "TENANT"
	RoleScopeUnit   RoleScopeType = "UNIT"
)

func NewRoleScopeType(v string) (RoleScopeType, *aerrs.AppError) {
	scope := RoleScopeType(v)

	if !scope.IsValid() {
		return "", aerrs.New(core.VALIDATION_FAILED,
			aerrs.WithAppendErrorDetail(
				aerrs.NewDetail(core.REASON_VAL_INVALID_FORMAT,
					aerrs.WithField("RoleScopeType"),
					aerrs.WithMessageDetail("invalid role scope type"),
				),
			),
		)
	}

	return scope, nil
}

func (r RoleScopeType) IsValid() bool {
	switch r {
	case RoleScopeGobal, RoleScopeTenant, RoleScopeUnit:
		return true
	default:
		return false
	}
}
