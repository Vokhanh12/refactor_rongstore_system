package entities

import (
	"github.com/google/uuid"

	cren "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/entities"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/validator"
	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

// ============================================================
// ENTITY
// ============================================================

type Role struct {
	cren.BaseEntity

	id uuid.UUID

	RoleKey vo.RoleKey
	name    string

	scopeType   enu.RoleScopeType
	accessScope enu.RoleAccessScope

	level int32

	description *string

	isSystem bool
	isSuper  bool
	isActive bool
}

// ============================================================
// PAYLOAD
// ============================================================

type RolePayload struct {
	RoleKey         vo.RoleKey
	Name            string
	RoleScopeType   enu.RoleScopeType
	RoleAccessScope enu.RoleAccessScope
	Level           int32
	Description     *string
	IsSystem        bool
	IsSuper         bool
	IsActive        bool
}

// ============================================================
// CONSTRUCTOR
// ============================================================

func NewRole(
	payload RolePayload,
) (*Role, *aerrs.AppError) {

	if err := validateRolePayload(payload); err != nil {
		return nil, err
	}

	role := &Role{
		id: uuid.Must(uuid.NewV7()),

		RoleKey: payload.RoleKey,
		name:    payload.Name,

		scopeType:   payload.RoleScopeType,
		accessScope: payload.RoleAccessScope,

		level: payload.Level,

		description: payload.Description,

		isSystem: payload.IsSystem,
		isSuper:  payload.IsSuper,
		isActive: payload.IsActive,
	}

	return role, nil
}

// ============================================================
// RESTORE
// ============================================================

func RestoreRole(
	id uuid.UUID,
	payload RolePayload,
) Role {

	return Role{
		id: id,

		RoleKey: payload.RoleKey,
		name:    payload.Name,

		scopeType:   payload.RoleScopeType,
		accessScope: payload.RoleAccessScope,

		level: payload.Level,

		description: payload.Description,

		isSystem: payload.IsSystem,
		isSuper:  payload.IsSuper,
		isActive: payload.IsActive,
	}
}

// ============================================================
// VALIDATION
// ============================================================

func validateRolePayload(
	payload RolePayload,
) *aerrs.AppError {

	v := validator.New()

	return v.
		Required("name", payload.Name).
		MaxLen("name", payload.Name, 100).
		RangeInt("level", int(payload.Level), 1, 255).
		Err()
}

// ============================================================
// DOMAIN METHODS
// ============================================================

func (r *Role) Activate() {
	r.isActive = true
}

func (r *Role) Deactivate() {

	if r.isSystem {
		return
	}

	r.isActive = false
}

// func (r *Role) PromoteLevel(
// 	level int32,
// ) *aerrs.AppError {

// 	if level < r.level {
// 		return aerrs.New("ROLE_LEVEL_INVALID")
// 	}

// 	r.level = level

// 	return nil
// }

func (r Role) IsElevated() bool {
	return r.isSuper
}

// ============================================================
// GETTERS
// ============================================================

func (r Role) ID() uuid.UUID {
	return r.id
}

func (r Role) RoleKey() vo.RoleKey {
	return r.RoleKey
}

func (r Role) Name() string {
	return r.name
}

func (r Role) ScopeType() enu.RoleScopeType {
	return r.scopeType
}

func (r Role) AccessScope() enu.RoleAccessScope {
	return r.accessScope
}

func (r Role) Level() int32 {
	return r.level
}

func (r Role) Description() *string {
	return r.description
}

func (r Role) IsSystem() bool {
	return r.isSystem
}

func (r Role) IsSuper() bool {
	return r.isSuper
}

func (r Role) IsActive() bool {
	return r.isActive
}
