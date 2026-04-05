package entities

import "time"

type RolePermission struct {
	Role       Role
	Permission Permission

	GrantedAt time.Time
	GrantedBy *string
}

func NewRolePermission(role Role, perm Permission) *RolePermission {
	return &RolePermission{
		Role:       role,
		Permission: perm,
	}
}
