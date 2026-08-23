package serializer

import (
	"encoding/json"

	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

func Marshal(v any) ([]byte, *aerrs.AppError) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, aerrs.New(
			core.JSON_SERIALIZATION_FAILED,
			aerrs.WithCauseDetail(err),
		)
	}

	return data, nil
}

func Unmarshal(data []byte, v any) *aerrs.AppError {
	if err := json.Unmarshal(data, v); err != nil {
		return aerrs.New(
			core.JSON_DESERIALIZATION_FAILED,
			aerrs.WithCauseDetail(err),
		)
	}

	return nil
}
