package handler

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/grpc"
)

var _ grpc.AuthzPort = (*AuthzHandler)(nil)

type AuthzHandler struct{}

// Handshake implements [ports.AuthzHandler].
func (g *AuthzHandler) Handshake(ctx context.Context, clientKey string) (session string, err error) {
	panic("unimplemented")
}

// Login implements [ports.AuthzHandler].
func (g *AuthzHandler) Login(ctx context.Context, username string, password string) (token string, err error) {
	panic("unimplemented")
}

func NewAuthzHandler() *AuthzHandler {
	return &AuthzHandler{}
}
