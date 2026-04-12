package apperrors

import commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"

func ToInternalError(err *AppError) *commonv1.Error {
	if err == nil {
		return nil
	}
	return &commonv1.Error{
		Code:         err.Code,
		Key:          err.Key,
		Message:      err.Message,
		Severity:     err.Severity,
		Retryable:    err.Retryable,
		HttpStatus:   int32(err.Status),
		GrpcCode:     err.GRPCCode,
		Source:       err.Source,
		ClientAction: err.ClientAction,
		ServerAction: err.ServerAction,
		Details:      MapAppErrorDetailsToProto(err.GetErrorDetails()),
	}
}

func ToPublicError(err *AppError) *commonv1.Error {
	if err == nil {
		return nil
	}
	return &commonv1.Error{
		Message:    err.Message,
		HttpStatus: int32(err.Status),
		GrpcCode:   err.GRPCCode,
		Details:    MapAppErrorDetailsToProto(err.GetErrorDetails()),
	}
}

func MapAppErrorDetailsToProto(
	details []AppErrorDetail,
) []*commonv1.ErrorDetail {

	if len(details) == 0 {
		return nil
	}

	result := make([]*commonv1.ErrorDetail, 0, len(details))
	for _, d := range details {
		result = append(result, &commonv1.ErrorDetail{
			Code:    d.Code,
			Field:   d.Field,
			Message: d.Message,
		})
	}
	return result
}
