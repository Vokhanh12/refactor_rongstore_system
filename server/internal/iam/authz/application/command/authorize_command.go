package command

import vo "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/authz/domain/valueobjects"

type AuthorizeCommand struct {
	UserID   string
	TenantID string
	RoleRef  []vo.RoleRef

	Resource   string
	Action     string
	ResourceID string
}

type AuthorizeCommandResult struct {
	Allowed bool
}
