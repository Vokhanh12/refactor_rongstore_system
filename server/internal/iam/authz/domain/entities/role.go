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

	id      uuid.UUID
	roleRef vo.RoleRef
	name    string

	scopeType   enu.RoleScopeType
	accessScope enu.RoleAccessScope
	level       int32
	description *string

	isSystem bool
	isSuper  bool
	isActive bool
}

// ============================================================
// PAYLOADS
// ============================================================

type RolePayload struct {
	RoleRef         vo.RoleRef
	RoleScopeType   enu.RoleScopeType
	Name            string
	RoleAccessScope enu.RoleAccessScope
	Level           int32
	Description     *string
	IsSystem        bool
	IsSuper         bool
	IsActive        bool
}

type NewRoleParams struct {
	RolePayload
}

type RestoreRoleParams struct {
	ID uuid.UUID
	RolePayload
}

// ============================================================
// CONSTRUCTOR (domain - validate)
// ============================================================

func NewRole(
	it NewRoleParams,
) (*Role, *aerrs.AppError) {

	err := validateRolePayload(it.RolePayload)

	if err != nil {
		return nil, err
	}

	role := newRoleFromPayload(
		uuid.Must(uuid.NewV7()),
		it.RolePayload,
	)

	return &role, nil
}

// ============================================================
// RESTORE (persistence - trust data)
// ============================================================

func RestoreRole(
	it RestoreRoleParams,
) Role {

	return newRoleFromPayload(
		it.ID,
		it.RolePayload,
	)
}

// ============================================================
// PRIVATE FACTORY
// ============================================================

func newRoleFromPayload(
	id uuid.UUID,
	payload RolePayload,
) Role {

	return Role{
		id:          id,
		roleRef:     payload.RoleRef,
		name:        payload.Name,
		scopeType:   payload.RoleScopeType,
		accessScope: payload.RoleAccessScope,
		level:       payload.Level,
		description: payload.Description,
		isSystem:    payload.IsSystem,
		isSuper:     payload.IsSuper,
		isActive:    payload.IsActive,
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
		Required("Name", payload.Name).
		MaxLen("Name", payload.Name, 100).
		RangeInt("Level", int(payload.Level), 1, 255).
		Err()
}

// ============================================================
// DOMAIN METHODS
// ============================================================

func (r Role) IsElevated() bool {
	return r.isSuper
}

// ============================================================
// GETTERS
// ============================================================

func (r Role) ID() uuid.UUID {
	return r.id
}

func (r Role) RoleRef() vo.RoleRef {
	return r.roleRef
}

func (r Role) Name() string {
	return r.name
}

func (r Role) RoleScopeType() enu.RoleScopeType {
	return r.scopeType
}

func (r Role) RoleAccessScope() enu.RoleAccessScope {
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
