package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

func DecodePayload(encoded string) (*Payload, error) {
	if encoded == "" {
		return nil, errors.New("empty jwt payload header")
	}

	raw, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	var payload Payload
	if err := json.Unmarshal(raw, &payload); err != nil {
		return nil, err
	}

	return &payload, nil
}
