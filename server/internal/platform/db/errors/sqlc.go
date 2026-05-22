package errors

import (
	aerr "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
)

type BaseRepository struct {
	DbError DBError
}

func (b *BaseRepository) HandleError(err error) *aerr.AppError {
	if err != nil {
		return TranslateDBError(err, b.DbError)
	}
	return nil
}
