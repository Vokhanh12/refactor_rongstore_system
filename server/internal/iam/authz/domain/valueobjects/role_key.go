package valueobjects

import (
	"strings"

	"github.com/google/uuid"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/infra/normalize"
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
// CONSTRUCTOR (domain - validate)
// ============================================================

func NewRoleKey(
	scopeID *uuid.UUID,
	roleCode string,
) (RoleKey, *aerrs.AppError) {

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

func RestoreRoleKey(
	roleCode string,
	scopeID *uuid.UUID,
) RoleKey {

	return RoleKey{
		roleCode: roleCode,
		scopeID:  scopeID,
	}
}

// ============================================================
// PARSER
// ============================================================

func ParseRoleKey(value string) (RoleKey, *aerrs.AppError) {

	parts := strings.Split(value, ":")

	v := validator.New().
		Format("roleKey", len(parts) == 2)

	if err := v.Err(); err != nil {
		return RoleKey{}, err
	}

	scopeID, err := normalize.ParseUUID(&parts[1])
	if err != nil {
		return RoleKey{}, err
	}

	return NewRoleKey(scopeID, parts[0])
}

// ============================================================
// GETTERS
// ============================================================

func (r RoleKey) RoleCode() string {
	return r.roleCode
}

func (r RoleKey) ScopeID() *uuid.UUID {
	return r.scopeID
}

// ============================================================
// UTILS
// ============================================================

func (r RoleKey) String() string {

	if r.scopeID == nil {
		return r.roleCode + ":<nil>"
	}

	return r.roleCode + ":" + r.scopeID.String()
}
