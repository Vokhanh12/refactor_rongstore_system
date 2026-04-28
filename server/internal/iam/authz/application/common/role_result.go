package common

import (
	"time"
)

type RoleResult struct {
	Id          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
