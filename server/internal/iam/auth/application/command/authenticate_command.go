package command

import "github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"

type AuthenticateCommand struct {
	Payload string
}

type AuthenticateCommandResult struct {
	Allowed  bool
	Identity ctxutil.IdentityContext
}
