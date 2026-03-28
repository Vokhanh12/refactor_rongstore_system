package entities

type Permission struct {
	ID   string
	Code string

	Name        *string
	Description *string

	Resource string
	Action   string

	IsActive bool
}
