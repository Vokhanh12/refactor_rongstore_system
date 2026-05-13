package valueobjects

import (
	"github.com/google/uuid"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/validator"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

// ============================================================
// VALUE OBJECT
// ============================================================

type RoleKey struct {
	roleCode string
	scopeID  *uuid.UUID
}

// ============================================================
// CONSTRUCTOR (domain - có validate)
// ============================================================

func NewRoleKey(scopeID *uuid.UUID, roleCode string) (RoleKey, *aerrs.AppError) {

	v := validator.New().
		Required("roleCode", roleCode)

	if err := v.Err(); err != nil {
		return RoleKey{}, err
	}

	return RoleKey{
		roleCode: roleCode,
		scopeID:  scopeID,
	}, nil
}

// ============================================================
// RESTORE (persistence - trust data)
// ============================================================

func RestoreRoleKey(roleCode string, scopeID *uuid.UUID) RoleKey {
	return RoleKey{
		roleCode: roleCode,
		scopeID:  scopeID,
	}
}

// ============================================================
// GETTERS
// ============================================================

func (r RoleKey) RoleCode() string    { return r.roleCode }
func (r RoleKey) ScopeID() *uuid.UUID { return r.scopeID }

// ============================================================
// UTILS
// ============================================================

func (r RoleKey) String() string {
	if r.scopeID == nil {
		return r.roleCode + ":<nil>"
	}
	return r.roleCode + ":" + r.scopeID.String()
}
