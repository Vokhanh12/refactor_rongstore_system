package infra

import "github.com/google/uuid"

func UUIDParse(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)

	if err != nil {
		return nil, aerrs.New(core.UUID_INVALID, aerrs.WithCauseDetail(err))
	}

	return id, nil
}
