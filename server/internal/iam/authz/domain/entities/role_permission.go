package entities

import (
	"time"

	"github.com/google/uuid"

	cren "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/domain/validator"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

// ============================================================
// ENTITY
// ============================================================

type RolePermission struct {
	cren.BaseEntity

	roleID       uuid.UUID
	permissionID uuid.UUID

	grantedAt time.Time
	grantedBy *uuid.UUID
}

// ============================================================
// PAYLOAD
// ============================================================

type RolePermissionPayload struct {
	RoleID       uuid.UUID
	PermissionID uuid.UUID

	GrantedAt time.Time
	GrantedBy *uuid.UUID
}

// ============================================================
// CONSTRUCTOR
// ============================================================

func NewRolePermission(
	payload RolePermissionPayload,
) (*RolePermission, *aerrs.AppError) {

	if err := validateRolePermissionPayload(payload); err != nil {
		return nil, err
	}

	rolePermission := &RolePermission{
		roleID:       payload.RoleID,
		permissionID: payload.PermissionID,

		grantedAt: payload.GrantedAt,
		grantedBy: payload.GrantedBy,
	}

	return rolePermission, nil
}

// ============================================================
// RESTORE
// ============================================================

func RestoreRolePermission(
	payload RolePermissionPayload,
) *RolePermission {

	return &RolePermission{
		roleID:       payload.RoleID,
		permissionID: payload.PermissionID,

		grantedAt: payload.GrantedAt,
		grantedBy: payload.GrantedBy,
	}
}

// ============================================================
// VALIDATION
// ============================================================

func validateRolePermissionPayload(
	payload RolePermissionPayload,
) *aerrs.AppError {

	v := validator.New()

	return v.
		Required("role_id", payload.RoleID.String()).
		Required("permission_id", payload.PermissionID.String()).
		Err()
}

// ============================================================
// GETTERS
// ============================================================

func (r RolePermission) RoleID() uuid.UUID {
	return r.roleID
}

func (r RolePermission) PermissionID() uuid.UUID {
	return r.permissionID
}

func (r RolePermission) GrantedAt() time.Time {
	return r.grantedAt
}

func (r RolePermission) GrantedBy() *uuid.UUID {
	return r.grantedBy
}
