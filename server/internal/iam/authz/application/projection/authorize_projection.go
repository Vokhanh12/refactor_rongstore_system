package projection

type AuthorizationGrant struct {
	RoleRef vo.RoleRef

	IsElevated bool

	Resource string
	Action   string
}

func (g AuthorizationGrant) Allow(resource, action string) bool {
	return g.Resource == resource && g.Action == action
}
