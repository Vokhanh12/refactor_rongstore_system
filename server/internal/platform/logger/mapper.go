package logger

import aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

func FromAppError(err *aerrs.AppError) LogEntry {
	if err == nil {
		return LogEntry{
			Message: "nil error",
		}
	}

	var causeDetail string
	if err.GetCauseDetail() != nil {
		causeDetail = err.GetCauseDetail().Error()
	}

	return LogEntry{
		Code:         err.Code,
		Key:          err.Key,
		Message:      err.Message,
		Cause:        err.Cause,
		CauseDetail:  causeDetail,
		ClientAction: err.ClientAction,
		ServerAction: err.ServerAction,
		Expected:     err.Expected,
		HTTPStatus:   err.Status,
		GRPCCode:     err.GRPCCode,
	}
}
