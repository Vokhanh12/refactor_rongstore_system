package projection

import vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"

type AuthorizationGrant struct {
	RoleKey vo.RoleKey

	IsElevated bool

	Resource string
	Action   string
}

func (g AuthorizationGrant) Match(resource, action string) bool {
	return g.Resource == resource && g.Action == action
}
