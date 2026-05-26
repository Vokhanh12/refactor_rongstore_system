package usecases

import (
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	platform "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

func Translate(err *aerrs.AppError) *aerrs.AppError {
	if err == nil {
		return nil
	}

	switch err.Code {

	case platform.DB_CONFLICT.Code:
		return wrapOperationFailure(err, core.REASON_APP_CONFLICT)

	case platform.DB_NOT_FOUND.Code:
		return wrapOperationFailure(err, core.REASON_APP_NOT_FOUND)

	case platform.DB_INVALID_REFERENCE.Code:
		return wrapOperationFailure(err, core.REASON_APP_DEPENDENCY_MISSING)
	}

	return err
}

func wrapOperationFailure(
	err *aerrs.AppError,
	reason aerrs.AppErrorDetail,
) *aerrs.AppError {
	return aerrs.New(
		core.APP_OPERATION_FAILED,
		aerrs.WithAppendErrorDetail(reason),
		aerrs.WithCauseDetail(err),
	)
}
