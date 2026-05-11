package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	errs "github.com/vokhanh12/refactor-rongstore-system/server/internal/iam/auth/errors"
	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

func DecodePayload(encoded string) (*Payload, *aerr.AppError) {

	if encoded == "" {
		return nil,
			aerr.New(
				errs.JWT_INVALID,
				aerr.WithCauseDetail(
					errors.New("empty jwt payload"),
				),
			)
	}

	raw, err := base64.RawURLEncoding.DecodeString(encoded)

	if err != nil {
		return nil,
			aerr.New(
				errs.JWT_PAYLOAD_INVALID,
				aerr.WithCauseDetail(err),
			)
	}

	var payload Payload

	if err := json.Unmarshal(raw, &payload); err != nil {

		return nil,
			aerr.New(
				errs.JWT_PAYLOAD_INVALID,
				aerr.WithCauseDetail(err),
			)
	}

	return &payload, nil
}
