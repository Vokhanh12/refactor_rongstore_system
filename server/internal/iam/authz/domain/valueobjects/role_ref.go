package valueobjects

import (
	"github.com/google/uuid"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

// ============================================================
// VALUE OBJECT
// ============================================================

type RoleRef struct {
	roleCode string
	scopeID  *uuid.UUID
}

// ============================================================
// CONSTRUCTOR (domain - có validate)
// ============================================================

func NewRoleRef(scopeID *uuid.UUID, roleCode string) (RoleRef, *aerrs.AppError) {

	r := RoleRef{
		roleCode: roleCode,
		scopeID:  scopeID,
	}

	err := r.validate()
	if err != nil {
		return RoleRef{}, err
	}

	return r, nil
}

// ============================================================
// RESTORE (persistence - trust data)
// ============================================================

func RestoreRoleRef(roleCode string, scopeID *uuid.UUID) RoleRef {
	return RoleRef{
		roleCode: roleCode,
		scopeID:  scopeID,
	}
}

// ============================================================
// VALIDATION
// ============================================================

func (r *RoleRef) validate() *aerrs.AppError {
	var details []aerrs.AppErrorDetail

	if r.roleCode == "" {
		details = append(details, aerrs.NewDetail(
			core.REASON_VAL_REQUIRED,
			aerrs.WithField("roleCode"),
			aerrs.WithMessageDetail("role code is required"),
		))
	}

	return aerrs.New(core.VALIDATION_FAILED, aerrs.WithAppendErrorDetails(details))
}

// ============================================================
// GETTERS
// ============================================================

func (r RoleRef) RoleCode() string    { return r.roleCode }
func (r RoleRef) ScopeID() *uuid.UUID { return r.scopeID }

// ============================================================
// UTILS
// ============================================================

func (r RoleRef) String() string {
	if r.scopeID == nil {
		return r.roleCode + ":<nil>"
	}
	return r.roleCode + ":" + r.scopeID.String()
}
