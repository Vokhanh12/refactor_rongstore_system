package grpc

import (
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

func ToGRPCError(grpcCode string, msg string) error {

	st := status.New(
		toCode(grpcCode),
		msg,
	)

	// if len(appErr.ErrorDetails) > 0 {
	// 	protoDetails := mapDetailsToProto(appErr.ErrorDetails)

	// 	stWithDetails, err := st.WithDetails(protoDetails...)
	// 	if err == nil {
	// 		st = stWithDetails
	// 	}
	// }

	return st.Err()
}

func toCode(code string) codes.Code {
	if c, ok := grpcCodeMap[code]; ok {
		return c
	}
	return codes.Unknown
}

// func mapDetailToProto(d aerrs.AppErrorDetail) *protos.ErrorDetail {
// 	return &protos.ErrorDetail{
// 		Field:   d.Field,
// 		Code:    d.Code,
// 		Message: d.Message,
// 		Hint:    d.Hint,
// 	}
// }

// func mapDetailsToProto(details []aerrs.AppErrorDetail) []interface{} {
// 	if len(details) == 0 {
// 		return nil
// 	}

// 	result := make([]interface{}, 0, len(details))
// 	for _, d := range details {
// 		result = append(result, mapDetailToProto(d))
// 	}
// 	return result
// }
