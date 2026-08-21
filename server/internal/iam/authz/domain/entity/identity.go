package entities

import vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"

type Identity struct {
	UserID       string
	RoleScopes   []vo.RoleScope
	AuthzVersion int
}
