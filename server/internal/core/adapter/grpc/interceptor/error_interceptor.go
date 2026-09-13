package interceptor

import (
	"context"
	"errors"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/adapter/mapper"
	dp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/dispatcher"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

	comv1rs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func ErrorUnaryInterceptor(
	logger Logger,
) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		resp, err := handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		return translateError(resp, err)
	}
}

// ============================================================
// Error translation
// ============================================================

func translateError(
	resp any,
	err error,
) (any, error) {
	if err == nil {
		return resp, nil
	}

	// Dispatcher collected multiple operation errors.
	var multipleErr *dp.MultipleError
	if errors.As(err, &multipleErr) {
		if _, ok := resp.(*comv1rs.MutateResponse); ok {
			return translateMutateDispatcherErrors(resp, multipleErr)
		}

		return translateViewDispatcherErrors(resp, multipleErr)
	}

	// Application/domain error.
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		return translateAppError(resp, appErr)
	}

	// Error is already a gRPC status.
	if status.Code(err) != codes.Unknown {
		return resp, err
	}

	// Unexpected error.
	return resp, status.Error(
		codes.Internal,
		"internal server error",
	)
}

// ============================================================
// Application error
// ============================================================

func translateAppError(
	resp any,
	appErr *apperrors.AppError,
) (any, error) {
	if appErr == nil {
		return resp, status.Error(
			codes.Internal,
			"internal server error",
		)
	}

	st := status.New(
		mapper.ToGRPCCode(appErr.GRPCCode),
		appErr.Message,
	)

	st, err := st.WithDetails(
		mapper.ToAppErrorInfo(appErr),
	)
	if err != nil {
		return resp, status.Error(
			codes.Internal,
			"internal server error",
		)
	}

	return resp, st.Err()
}

// ============================================================
// Dispatcher multiple errors
// ============================================================

func translateDispatcherErrors(
	resp any,
	multipleErr *dp.MultipleError,
) (any, error) {

	if multipleErr == nil || multipleErr.Empty() {
		return resp, nil
	}

	st := status.New(200, "OK")

	details := make([]proto.Message, 0)

	for _, opErr := range multipleErr.Errors {
		if opErr == nil {
			continue
		}

		var appErr *apperrors.AppError

		if errors.As(opErr.Err, &appErr) {
			details = append(
				details,
				mapper.ToAppErrorInfo(opErr.OpID, appErr),
			)

			continue
		}

		details = append(
			details, mapper.ToAppErrorInfo(opErr.OpID, &apperrors.INTERNAL_FALLBACK))

	}

	details = append(
		details,
		mapper.ToDispatchSummary(multipleErr),
	)

	return resp, st.Err()
}
