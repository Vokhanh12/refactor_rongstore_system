package entities

import (
	"time"

	"github.com/google/uuid"

	cren "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/domain/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/domain/validator"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

// ============================================================
// ENUM
// ============================================================

type ScopeType string

const (
	ScopeGlobal       ScopeType = "GLOBAL"
	ScopeOrganization ScopeType = "ORGANIZATION"
	ScopeUnit         ScopeType = "UNIT"
)

// ============================================================
// ENTITY
// ============================================================

type UserRole struct {
	cren.BaseEntity

	userID uuid.UUID
	roleID uuid.UUID

	scopeType ScopeType
	scopeID   *string

	assignedAt time.Time
	assignedBy *uuid.UUID
}

// ============================================================
// PAYLOAD
// ============================================================

type UserRolePayload struct {
	UserID uuid.UUID
	RoleID uuid.UUID

	ScopeType ScopeType
	ScopeID   *string

	AssignedAt time.Time
	AssignedBy *uuid.UUID
}

// ============================================================
// CONSTRUCTOR
// ============================================================

func NewUserRole(
	payload UserRolePayload,
) (*UserRole, *aerrs.AppError) {

	if err := validateUserRolePayload(payload); err != nil {
		return nil, err
	}

	userRole := &UserRole{
		userID: payload.UserID,
		roleID: payload.RoleID,

		scopeType: payload.ScopeType,
		scopeID:   payload.ScopeID,

		assignedAt: payload.AssignedAt,
		assignedBy: payload.AssignedBy,
	}

	return userRole, nil
}

// ============================================================
// RESTORE
// ============================================================

func RestoreUserRole(
	payload UserRolePayload,
) *UserRole {

	return &UserRole{
		userID: payload.UserID,
		roleID: payload.RoleID,

		scopeType: payload.ScopeType,
		scopeID:   payload.ScopeID,

		assignedAt: payload.AssignedAt,
		assignedBy: payload.AssignedBy,
	}
}

// ============================================================
// VALIDATION
// ============================================================

// ============================================================
// VALIDATION
// ============================================================

func validateUserRolePayload(
	payload UserRolePayload,
) *aerrs.AppError {

	v := validator.New()

	v.Required("user_id", payload.UserID.String())
	v.Required("role_id", payload.RoleID.String())

	validScope := payload.ScopeType.IsValid()

	v.Enum("scope_type", validScope)

	if !validScope {
		return v.Err()
	}

	switch payload.ScopeType {

	case ScopeGlobal:

	case ScopeOrganization, ScopeUnit:
		if payload.ScopeID == nil || *payload.ScopeID == "" {
			v.Required("scope_id", "")
		}
	}

	return v.Err()
}

// ============================================================
// DOMAIN METHODS
// ============================================================

func (u UserRole) IsGlobalScope() bool {
	return u.scopeType == ScopeGlobal
}

func (u UserRole) IsOrganizationScope() bool {
	return u.scopeType == ScopeOrganization
}

func (u UserRole) IsUnitScope() bool {
	return u.scopeType == ScopeUnit
}

func (s ScopeType) IsValid() bool {
	switch s {
	case ScopeGlobal, ScopeOrganization, ScopeUnit:
		return true
	default:
		return false
	}
}

// ============================================================
// GETTERS
// ============================================================

func (u UserRole) UserID() uuid.UUID {
	return u.userID
}

func (u UserRole) RoleID() uuid.UUID {
	return u.roleID
}

func (u UserRole) ScopeType() ScopeType {
	return u.scopeType
}

func (u UserRole) ScopeID() *string {
	return u.scopeID
}

func (u UserRole) AssignedAt() time.Time {
	return u.assignedAt
}

func (u UserRole) AssignedBy() *uuid.UUID {
	return u.assignedBy
}
