package apperrors

import (
	proto "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
		detail := &proto.ErrorDetail{
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
) []*proto.ErrorDetail {

	if len(details) == 0 {
		return nil
	}

	result := make([]*proto.ErrorDetail, 0, len(details))
	for _, d := range details {
		result = append(result, &proto.ErrorDetail{
			Code:    d.Code,
			Field:   d.Field,
			Message: d.Message,
		})
	}
	return result
}
