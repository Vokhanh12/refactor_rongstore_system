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

func ToGRPCError(code string, msg string) error {

	st := status.New(
		toGRPCCode(code),
		msg,
	)

	// // attach details (batch)
	// if len(appErr.ErrorDetails) > 0 {
	// 	protoDetails := mapDetailsToProto(appErr.ErrorDetails)

	// 	stWithDetails, err := st.WithDetails(protoDetails...)
	// 	if err == nil {
	// 		st = stWithDetails
	// 	}
	// 	// nếu fail thì ignore (không phá main error)
	// }

	return st.Err()
}

func toGRPCCode(code string) codes.Code {
	if c, ok := grpcCodeMap[code]; ok {
		return c
	}
	return codes.Unknown
}
