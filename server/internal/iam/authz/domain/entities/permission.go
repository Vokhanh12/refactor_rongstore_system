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

func (p Permission) Key() string {
	return p.Resource + ":" + p.Action
}

func (p Permission) Match(resource, action string) bool {
	return p.Resource == resource && p.Action == action
}
