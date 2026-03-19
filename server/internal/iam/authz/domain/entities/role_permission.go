package entities

import "time"

type RolePermission struct {
	RoleID       string
	PermissionID string

	GrantedAt time.Time
	GrantedBy *string
}
