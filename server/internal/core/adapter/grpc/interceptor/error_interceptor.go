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
	"google.golang.org/protobuf/protoadapt"
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

func translateError(resp any, err error) (any, error) {
	// Dispatcher multiple errors.
	var dispatcherErrors *dp.Errors
	if errors.As(err, &dispatcherErrors) {
		return translateDispatcherErrors(resp, dispatcherErrors)
	}

	// Application error.
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		return translateAppError(resp, appErr)
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

// ============================================================
// Application error
// ============================================================

func translateAppError(
	resp any,
	appErr *apperrors.AppError,
) (any, error) {
	st := status.New(
		mapper.ToGRPCCode(appErr.GRPCCode),
		appErr.Message,
	)

	errorInfo := &comv1rs.AppErrorInfo{
		Metadata: &comv1rs.MetadataReponse{
			OpId: "uuid",
		},
		Reason:     appErr.Code,
		Domain:     appErr.Domain,
		Layer:      appErr.Layer,
		Message:    appErr.Message,
		Violations: mapper.ToProtoViolations(appErr.Violations),
	}

	st, err := st.WithDetails(errorInfo)
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
	dispatcherErrors *dp.Errors,
) (any, error) {
	mutateResp, ok := resp.(*comv1rs.MutateResponse)
	if !ok {
		return resp, status.Error(
			codes.Internal,
			"invalid mutation response",
		)
	}

	errorInfos := make([]protoadapt.MessageV1, 0)

	for _, mutation := range mutateResp.MutateResults {
		if mutation.Error == nil {
			continue
		}

		errorInfos = append(
			errorInfos,
			protoadapt.MessageV1Of(mutation.Error),
		)
	}

	if len(errorInfos) == 0 {
		return resp, status.Error(
			codes.Internal,
			"mutation failed without error details",
		)
	}

	// Status code lấy từ error đầu tiên.
	st := status.New(
		mapper.ToGRPCCode(dispatcherErrors.Errors[0].Err.GRPCCode),
		dispatcherErrors.Errors[0].Err.Message,
	)

	st, err := st.WithDetails(errorInfos...)
	if err != nil {
		return resp, status.Error(
			codes.Internal,
			"internal server error",
		)
	}

	return resp, st.Err()
}
