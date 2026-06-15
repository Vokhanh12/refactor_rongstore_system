package usecases

import (
	core "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	platform "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

func Translate(repoer *aerrs.AppError) *aerrs.AppError {
	if repoer == nil {
		return nil
	}

	switch repoer.Code {

	case platform.DB_CONFLICT.Code:
		return wrapOperationFailure(repoer, core.REASON_APP_CONFLICT)

	case platform.DB_NOT_FOUND.Code:
		return wrapOperationFailure(repoer, core.REASON_APP_NOT_FOUND)

	case platform.DB_INVALID_REFERENCE.Code:
		return wrapOperationFailure(repoer, core.REASON_APP_DEPENDENCY_MISSING)
	}

	return repoer
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
