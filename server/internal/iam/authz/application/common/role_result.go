package common

import (
	"time"

	"github.com/google/uuid"
)

type RoleResult struct {
	Id          uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
