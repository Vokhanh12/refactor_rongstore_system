package entities

import (
	"github.com/google/uuid"

	enu "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/enums"
	vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"
	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type Role struct {
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

func NewRole(vldRoleRef vo.ValidatedRoleRef, name string, roleScopeType enu.RoleScopeType,
	roleAccessScope enu.RoleAccessScope, level uint8, description *string, isSystem bool, isSuper bool, isActive bool) Role {

	return Role{
		id:          uuid.Must(uuid.NewV7()),
		roleRef:     vldRoleRef.RoleRef,
		name:        name,
		scopeType:   roleScopeType,
		accessScope: roleAccessScope,
		level:       level,
		description: description,
		isSystem:    isSystem,
		isSuper:     isSuper,
		isActive:    isActive,
	}
}

func (p *Role) validate() []aerrs.AppErrorDetail {
	var details []aerrs.AppErrorDetail

	if p.name == "" {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_REQUIRED,
			aerrs.WithField("name"),
		))
	}

	if p.level > 255 {
		details = append(details, *aerrs.NewDetail(
			domain.REASON_OUT_OF_RANGE,
			aerrs.WithField("level"),
		))
	}

	if len(details) > 0 {
		return details
	}

	return nil
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

func NewRoleFromPersistence(p NewRoleParams) Role {
	return Role{
		id:          p.ID,
		roleRef:     p.RoleRef,
		name:        p.Name,
		scopeType:   p.RoleScopeType,
		accessScope: p.RoleAccessScope,
		level:       p.Level,
		description: p.Description,
		isSystem:    p.IsSystem,
		isSuper:     p.IsSuper,
		isActive:    p.IsActive,
	}
}

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
