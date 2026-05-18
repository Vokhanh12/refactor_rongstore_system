package ports

import (
	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type Payload struct {
	UserID      string   `json:"sub"`
	RoleKeyStrs []string `json:"roles,omitempty"`
}

type TokenProvider interface {
	DecorePayload(encoded string) (*Payload, *aerr.AppError)
}
