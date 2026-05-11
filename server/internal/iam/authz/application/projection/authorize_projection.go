package projection

type AuthorizationGrant struct {
	RoleRef vo.RoleRef

	IsElevated bool

	Resource string
	Action   string
}
