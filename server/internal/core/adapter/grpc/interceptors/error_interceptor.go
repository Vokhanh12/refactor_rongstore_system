package interceptor

import (
	"context"
	"errors"

	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ErrorUnaryInterceptor(
	logger Logger,
) grpc.UnaryServerInterceptor {

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		resp, err := handler(ctx, req)

		if err == nil {
			return resp, nil
		}

		return translateRespError(resp, err)
	}
}

func translateRespError(resp any, err error) (any, error) {
	// Application error.
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		return ToGRPCError(resp, appErr)
	}

	// Already a gRPC status.
	if status.Code(err) != codes.Unknown {
		return resp, err
	}

	// Unknown/unexpected error.
	return resp, status.Error(
		codes.Internal,
		"internal server error",
	)
}

// ===>>> Điều chỉnh dispatcher result trả về nhiều error
func ToGRPCError(resp any, appErr *apperrors.AppError) (any, error) {
	st := status.New(
		toGRPCCode(appErr.GRPCCode),
		appErr.Message,
	)

	errorInfo := &errdetails.ErrorInfo{
		Reason: appErr.Code,
		Domain: appErr.Domain,
		Metadata: map[string]string{
			"layer": appErr.Layer,
		},
	}

	if len(appErr.Violations) == 0 {
		stWithDetails, err := st.WithDetails(errorInfo)
		if err != nil {
			return status.Error(
				codes.Internal,
				"internal server error",
			)
		}

		return stWithDetails.Err()
	}

	badRequest := &errdetails.BadRequest{}

	for _, violation := range appErr.Violations {
		badRequest.FieldViolations = append(
			badRequest.FieldViolations,
			&errdetails.BadRequest_FieldViolation{
				Reason:      violation.Code,
				Field:       violation.Field,
				Description: violation.Message,
			},
		)
	}

	stWithDetails, err := st.WithDetails(
		errorInfo,
		badRequest,
	)

	if err != nil {
		return status.Error(
			codes.Internal,
			"internal server error",
		)
	}

	return stWithDetails.Err()
}
