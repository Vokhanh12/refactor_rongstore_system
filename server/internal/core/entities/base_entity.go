package entities

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt time.Time
}

func NewBase(id uuid.UUID) BaseEntity {
	now := time.Now()
	return BaseEntity{
		id:        id,
		createdAt: now,
		updatedAt: now,
	}
}

func RestoreBase(id uuid.UUID, createdAt, updatedAt time.Time) BaseEntity {
	return BaseEntity{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

// getters only
func (b BaseEntity) ID() uuid.UUID        { return b.id }
func (b BaseEntity) CreatedAt() time.Time { return b.createdAt }
func (b BaseEntity) UpdatedAt() time.Time { return b.updatedAt }

// internal only
func (b *BaseEntity) touch() {
	b.updatedAt = time.Now()
}
