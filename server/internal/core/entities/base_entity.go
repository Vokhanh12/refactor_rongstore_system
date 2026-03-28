package entities

import "time"

type BaseEntity struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBaseEntity(id string) BaseEntity {
	return BaseEntity{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
