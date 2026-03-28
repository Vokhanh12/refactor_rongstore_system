package entities

import "time"

type SoftDeleteEntity struct {
	BaseEntity
	DeletedAt *time.Time
}

func (e *SoftDeleteEntity) Delete() {
	now := time.Now()
	e.DeletedAt = &now
}
