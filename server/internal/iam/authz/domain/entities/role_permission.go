package entities

import "time"

type RolePermission struct {
	Role       Role
	Permission Permission

	GrantedAt time.Time
	GrantedBy *string
}
