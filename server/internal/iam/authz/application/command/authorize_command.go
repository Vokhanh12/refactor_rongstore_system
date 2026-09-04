package command

import "github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"

type AuthorizeCommand struct {
	UserID string

	RoleScopes []ctxutil.RoleScope

	Resource   string
	Action     string
	ResourceID string
}

type AuthorizeCommandResult struct {
	Allowed bool
}
