package model

import "time"

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy *string
	UpdatedBy *string
}
