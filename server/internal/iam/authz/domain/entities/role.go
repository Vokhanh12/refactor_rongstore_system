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
	level       uint8
	description *string

	isSystem bool
	isSuper  bool
	isActive bool
}

type NewRoleParams struct {
	ID              uuid.UUID
	RoleRef         vo.RoleRef
	RoleScopeType   enu.RoleScopeType
	Name            string
	RoleAccessScope enu.RoleAccessScope
	Level           uint8
	Description     *string
	IsSystem        bool
	IsSuper         bool
	IsActive        bool
}

// ============================================================
// CONSTRUCTOR (domain - có validate)
// ============================================================

func NewRole(it NewRoleParams) (*Role, *aerrs.AppError) {

	v := validator.New()

	err := v.
		Required("Name", it.Name).
		MaxLen("Name", it.Name, 100).
		RangeInt("level", int(it.Level), 1, 255).
		Err()

	if err != nil {
		return nil, err
	}

	return &Role{
		id:          it.ID,
		roleRef:     it.RoleRef,
		name:        it.Name,
		scopeType:   it.RoleScopeType,
		accessScope: it.RoleAccessScope,
		level:       it.Level,
		description: it.Description,
		isSystem:    it.IsSystem,
		isSuper:     it.IsSuper,
		isActive:    it.IsActive,
	}, nil
}

// ============================================================
// RESTORE (persistence - trust data)
// ============================================================

func RestoreRole(it NewRoleParams) Role {
	return Role{
		id:          it.ID,
		roleRef:     it.RoleRef,
		name:        it.Name,
		scopeType:   it.RoleScopeType,
		accessScope: it.RoleAccessScope,
		level:       it.Level,
		description: it.Description,
		isSystem:    it.IsSystem,
		isSuper:     it.IsSuper,
		isActive:    it.IsActive,
	}
}

// ============================================================
// SETTERS (domain - validate)
// ============================================================

// ============================================================
// GETTERS
// ============================================================

func (r Role) IsElevated() bool {
	return r.isSuper
}

func (p Role) RoleRef() vo.RoleRef {
	return p.roleRef
}

func (r Role) ID() uuid.UUID {
	return r.id
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

func (r Role) Level() uint8 {
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
