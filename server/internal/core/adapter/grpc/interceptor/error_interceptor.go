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
		return translateDispatcherErrors(resp, multipleErr)
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
		toAppErrorInfo(appErr),
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

	/*
		MultipleError only contains failed operations.

		Therefore:
			- There is no success information here.
			- We must not calculate success count here.
			- The RPC is considered failed.
	*/

	firstErr := multipleErr.FirstError()
	if firstErr == nil {
		return resp, status.Error(
			codes.Internal,
			"internal server error",
		)
	}

	/*
		If the first error is an AppError, use it as the
		representative gRPC status.

		All other AppErrors are attached as status details.
	*/
	firstAppErr := multipleErr.FirstAppError()

	if firstAppErr == nil {
		return resp, status.Error(
			codes.Internal,
			"internal server error",
		)
	}

	st := status.New(
		mapper.ToGRPCCode(firstAppErr.GRPCCode),
		firstAppErr.Message,
	)

	details := make([]proto.Message, 0)

	for _, appErr := range multipleErr.AppErrors() {
		details = append(
			details,
			toAppErrorInfo(appErr),
		)
	}

	/*
		DispatchSummary is metadata describing the batch.

		It is optional, but if you want the client to know that
		multiple operations failed, it can be attached here.
	*/
	details = append(
		details,
		toDispatchSummary(multipleErr),
	)

	st, err := st.WithDetails(details...)
	if err != nil {
		return resp, status.Error(
			codes.Internal,
			"internal server error",
		)
	}

	return resp, st.Err()
}

// ============================================================
// AppError → Proto
// ============================================================

func toAppErrorInfo(
	appErr *apperrors.AppError,
) *comv1rs.AppErrorInfo {
	if appErr == nil {
		return nil
	}

	return &comv1rs.AppErrorInfo{
		Metadata: &comv1rs.MetadataReponse{
			OpId: appErr.OpID,
		},
		Reason:     appErr.Code,
		Domain:     appErr.Domain,
		Layer:      appErr.Layer,
		Message:    appErr.Message,
		Violations: mapper.ToProtoViolations(appErr.Violations),
	}
}

// ============================================================
// Dispatcher summary
// ============================================================

func toDispatchSummary(
	multipleErr *dp.MultipleError,
) *comv1rs.DispatchSummary {
	if multipleErr == nil {
		return nil
	}

	total := multipleErr.Total()

	return &comv1rs.DispatchSummary{
		Total:     int32(total),
		Succeeded: 0,
		Failed:    int32(total),
	}
}
