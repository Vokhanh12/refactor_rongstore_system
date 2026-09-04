package interceptor

import (
	"context"
	"errors"

	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
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

		return resp, translateError(err)
	}
}

func translateError(err error) error {
	var appErr *apperrors.AppError

	if errors.As(err, &appErr) {
		return ToGRPCError(
			appErr.Code,
			appErr.Message,
		)
	}

	// Đã là gRPC status thì giữ nguyên.
	if _, ok := status.FromError(err); ok {
		return err
	}

	// Unknown error.
	return status.Error(
		codes.Internal,
		"internal server error",
	)
}
