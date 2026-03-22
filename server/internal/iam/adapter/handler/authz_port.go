package handler

import "context"

type AuthzPort interface {
	Login(ctx context.Context, username, password string) (token string, err error)
	Handshake(ctx context.Context, clientKey string) (session string, err error)
}
