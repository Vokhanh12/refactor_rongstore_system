package errors

import (
	pkgerrors "server/pkg/errors"
)

type BaseRepository struct {
	DbError DBError
}

func (b *BaseRepository) HandleError(err error) *pkgerrors.AppError {
	if err != nil {
		return TranslateDBError(err, b.DbError)
	}
	return nil
}
