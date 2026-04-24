package valueobjects

import (
	"github.com/google/uuid"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/validator"
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

	v := validator.New().
		Required("roleCode", roleCode)

	if err := v.Err(); err != nil {
		return RoleRef{}, err
	}

	return RoleRef{
		roleCode: roleCode,
		scopeID:  scopeID,
	}, nil
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
