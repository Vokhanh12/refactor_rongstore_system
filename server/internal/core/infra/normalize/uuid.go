package normalize

import (
	"strings"

	"github.com/google/uuid"
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

func ParseUUID(s *string) (*uuid.UUID, *aerrs.AppError) {

	if s == nil {
		return nil, nil
	}

	v := strings.TrimSpace(*s)

	if v == "" {
		return nil, nil
	}

	id, err := uuid.Parse(v)

	if err != nil {
		return nil,
			aerrs.New(core.UUID_INVALID,
				aerrs.WithCauseDetail(err))
	}

	return &id, nil
}
