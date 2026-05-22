package serialization

import (
	"encoding/json"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	"github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

func MustMarshal(v any) ([]byte, *apperrors.AppError) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, aerrs.New(core.JSON_SERIALIZATION_FAILED)
	}
	return b, nil
}
