package projection

import vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"

type AuthorizationGrant struct {
	RoleKey vo.RoleKey

	IsElevated bool

	ResourceAction vo.ResourceAction
}
