package valueobjects

import "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"

type ValidatedRoleRef struct {
	RoleRef
	isValidated bool
}

func (vp *ValidatedRoleRef) IsValid() bool {
	return vp.isValidated
}

func NewValidatedRoleRef(RoleRef *RoleRef) (*ValidatedRoleRef, []apperrors.AppErrorDetail) {
	if err := RoleRef.validate(); err != nil {
		return nil, err
	}

	return &ValidatedRoleRef{
		RoleRef:     *RoleRef,
		isValidated: true,
	}, nil
}
