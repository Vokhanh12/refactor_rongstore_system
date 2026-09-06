package interceptor

import "google.golang.org/grpc/codes"

var grpcCodeMap = map[string]codes.Code{
	"OK":                 codes.OK,
	"Canceled":           codes.Canceled,
	"Unknown":            codes.Unknown,
	"InvalidArgument":    codes.InvalidArgument,
	"DeadlineExceeded":   codes.DeadlineExceeded,
	"NotFound":           codes.NotFound,
	"AlreadyExists":      codes.AlreadyExists,
	"PermissionDenied":   codes.PermissionDenied,
	"ResourceExhausted":  codes.ResourceExhausted,
	"FailedPrecondition": codes.FailedPrecondition,
	"Aborted":            codes.Aborted,
	"OutOfRange":         codes.OutOfRange,
	"Unimplemented":      codes.Unimplemented,
	"Internal":           codes.Internal,
	"Unavailable":        codes.Unavailable,
	"DataLoss":           codes.DataLoss,
	"Unauthenticated":    codes.Unauthenticated,
}

func toGRPCCode(code string) codes.Code {
	if c, ok := grpcCodeMap[code]; ok {
		return c
	}

	return codes.Internal
}
