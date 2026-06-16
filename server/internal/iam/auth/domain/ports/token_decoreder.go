package services

import (
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type Payload struct {
	UserID      string   `json:"sub"`
	RoleKeyStrs []string `json:"roles,omitempty"`
}

type TokenDecoreder interface {
	DecorePayload(encoded string) (*Payload, *aerrs.AppError)
}
