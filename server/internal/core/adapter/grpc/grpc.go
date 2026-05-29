package grpc

import (
	protos "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var grpcCodeMap = map[string]codes.Code{
	"Internal":         codes.Internal,
	"Unavailable":      codes.Unavailable,
	"InvalidArgument":  codes.InvalidArgument,
	"Unauthenticated":  codes.Unauthenticated,
	"AlreadyExists":    codes.AlreadyExists,
	"NotFound":         codes.NotFound,
	"PermissionDenied": codes.PermissionDenied,
	"DeadlineExceeded": codes.DeadlineExceeded,
}

func ToGRPCError(appErr *aerrs.AppError) error {
	if appErr == nil {
		return status.Error(codes.Internal, "internal error")
	}

	st := status.New(
		toGRPCCode(appErr.GRPCCode),
		appErr.Message,
	)

	// attach details (batch)
	if len(appErr.ErrorDetails) > 0 {
		protoDetails := mapDetailsToProto(appErr.ErrorDetails)

		stWithDetails, err := st.WithDetails(protoDetails...)
		if err == nil {
			st = stWithDetails
		}
		// nếu fail thì ignore (không phá main error)
	}

	return st.Err()
}

func toGRPCCode(code string) codes.Code {
	if c, ok := grpcCodeMap[code]; ok {
		return c
	}
	return codes.Unknown
}

func mapDetailToProto(d aerrs.AppErrorDetail) *protos.ErrorDetail {
	return &protos.ErrorDetail{
		Field:   d.Field,
		Code:    d.Code,
		Message: d.Message,
		Hint:    d.Hint,
	}
}

func mapDetailsToProto(details []aerrs.AppErrorDetail) []interface{} {
	if len(details) == 0 {
		return nil
	}

	result := make([]interface{}, 0, len(details))
	for _, d := range details {
		result = append(result, mapDetailToProto(d))
	}
	return result
}
