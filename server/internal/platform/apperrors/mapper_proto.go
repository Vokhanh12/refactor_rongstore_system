package apperrors

import (
	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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

func ToGRPCError(appErr *AppError) error {

	var grpcCode codes.Code

	switch appErr.GRPCCode {
	case "Internal":
		grpcCode = codes.Internal

	case "Unavailable":
		grpcCode = codes.Unavailable

	case "InvalidArgument":
		grpcCode = codes.InvalidArgument

	case "Unauthenticated":
		grpcCode = codes.Unauthenticated

	case "AlreadyExists":
		grpcCode = codes.AlreadyExists

	case "NotFound":
		grpcCode = codes.NotFound

	default:
		grpcCode = codes.Unknown
	}

	st := status.New(grpcCode, appErr.Message)

	for _, d := range appErr.GetErrorDetails() {
		detail := &commonv1.ErrorDetail{
			Field:   d.Field,
			Code:    d.Code,
			Message: d.Message,
			Hint:    d.Hint,
		}

		stWithDetails, err := st.WithDetails(detail)
		if err != nil {
			continue // không fail toàn bộ vì detail lỗi
		}
		st = stWithDetails
	}

	return st.Err()
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
