package handler

import (
	"context"

	authrs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/iam/auth/v1/resources"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// Handshake implements [ports.AuthHandler].
func (g *AuthHandler) Handshake(ctx context.Context, clientKey string) (session string, err error) {
	panic("unimplemented")
}

// Login implements [ports.AuthHandler].
func (g *AuthHandler) Login(ctx context.Context, req *authrs.RoleMutateRequest) (token string, err error) {
	panic("unimplemented")
}
