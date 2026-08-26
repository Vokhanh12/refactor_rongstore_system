package serializer

import (
	"encoding/json"
)

func Marshal(v any) ([]byte, error) {
	data, err := json.Marshal(v)
	if err != nil {
		wrap("marshal", err)
	}

	return data, nil
}

func Unmarshal(data []byte, v any) error {
	if err := json.Unmarshal(data, v); err != nil {
		wrap("unmarshal", err)
	}

	return nil
}
