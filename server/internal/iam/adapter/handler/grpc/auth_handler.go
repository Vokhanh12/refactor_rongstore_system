package handler

import (
	"context"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/adapter/handler"
)

var _ handler.AuthPort = (*AuthHandler)(nil)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// Handshake implements [ports.AuthHandler].
func (g *AuthHandler) Handshake(ctx context.Context, clientKey string) (session string, err error) {
	panic("unimplemented")
}

// Login implements [ports.AuthHandler].
func (g *AuthHandler) Login(ctx context.Context, username string, password string) (token string, err error) {
	panic("unimplemented")
}
