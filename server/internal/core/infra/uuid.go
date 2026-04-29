package infra

import (
	"github.com/google/uuid"
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

func UUIDParse(s string) (*uuid.UUID, *aerrs.AppError) {
	id, err := uuid.Parse(s)

	if err != nil {
		return nil, aerrs.New(core.UUID_INVALID, aerrs.WithCauseDetail(err))
	}

	return &id, nil
}
