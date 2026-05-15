package valueobjects

import "github.com/google/uuid"

type UUID struct {
	value string
}

func NewUUID() UUID {
	return UUID{value: uuid.NewString()}
}

func (u UUID) String() string {
	return u.value
}
