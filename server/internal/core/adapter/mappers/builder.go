package mappers

import (
	"context"
	"time"

	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func BuildMutateResult(opID string, data any, err *aerrs.AppError) dtos.MutateResultItem {
	return dtos.MutateResultItem{
		OpID:  opID,
		Data:  data,
		Error: AppErrorToDTO(err),
	}
}

func BuildTransportErrorResponse(ctx context.Context, err *aerrs.AppError) (*commonv1.BaseResponse, error) {

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	return &commonv1.BaseResponse{
		Success: false,
		Metadata: &commonv1.Metadata{
			TraceId:    requestctx.TraceID,
			RequestId:  requestctx.RequestID,
			Locale:     locatectx.Locale,
			Region:     locatectx.Region,
			Degraded:   false,
			ServerTime: time.Now().UnixMilli(),
		},
		Error: AppErrorToDTO(err),

		// &commonv1.Error{
		// 	//Code:         err.Code,
		// 	//Key:          err.Key,
		// 	Message: err.Message,
		// 	//Severity:     err.Severity,
		// 	//Retryable:    err.Retryable,
		// 	//Source:       err.Source,
		// 	GrpcCode:   err.GRPCCode,
		// 	HttpStatus: int32(err.Status),
		// 	//ClientAction: err.ClientAction,
		// 	//ServerAction: err.ServerAction,
		// 	Details: aerrs.MapAppErrorDetailsToProto(err.GetErrorDetails()),
		// },
	}, nil

}

func BuildMutateErrorResponse(ctx context.Context, errs []aerrs.AppError) *commonv1.MutateResponse {

	results := make([]*commonv1.MutateResult, 0, len(errs))

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	for _, err := range errs {
		results = append(results, &commonv1.MutateResult{
			Success: false,
			Error: &commonv1.Error{
				Code:         err.Code,
				Key:          err.Key,
				Message:      err.Message,
				Severity:     err.Severity,
				Retryable:    err.Retryable,
				Source:       err.Source,
				GrpcCode:     err.GRPCCode,
				HttpStatus:   int32(err.Status),
				ClientAction: err.ClientAction,
				ServerAction: err.ServerAction,
				Details:      aerrs.MapAppErrorDetailsToProto(err.GetErrorDetails()),
			},
		})
	}

	return &commonv1.MutateResponse{
		Metadata: &commonv1.Metadata{
			TraceId:    requestctx.TraceID,
			RequestId:  requestctx.RequestID,
			Locale:     locatectx.Locale,
			Region:     locatectx.Region,
			Degraded:   false,
			ServerTime: time.Now().UnixMilli(),
		},
		MutateResults: results,
	}
}

// BuildSuccessResponse builds a BaseResponse representing a successful operation.
// It converts the provided protobuf message to google.protobuf.Any.
func BuildSuccessResponse(ctx context.Context, data proto.Message) (*commonv1.BaseResponse, error) {
	anyData, err := anypb.New(data)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to marshal response")
	}

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	return &commonv1.BaseResponse{
		Success: true,
		Data:    anyData,
		Metadata: &commonv1.Metadata{
			TraceId:    requestctx.TraceID,
			RequestId:  requestctx.RequestID,
			Locale:     locatectx.Locale,
			Region:     locatectx.Region,
			Degraded:   false,
			ServerTime: time.Now().UnixMilli(),
		},
	}, nil
}

// // FromError converts a generic error into a BaseResponse using BusinessError mapping.
// func FromError(ctx context.Context, err error) *commonv1.BaseResponse {
// 	be, _ := aerrs.GetBusinessError(err)
// 	return BuildErrorResponse(ctx, be)
// }

func BuildMutateResponse(
	ctx context.Context,
	items []*commonv1.MutateResult,
) *commonv1.MutateResponse {

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	return &commonv1.MutateResponse{
		Metadata: &commonv1.Metadata{
			TraceId:    requestctx.TraceID,
			RequestId:  requestctx.RequestID,
			Locale:     locatectx.Locale,
			Region:     locatectx.Region,
			Degraded:   false,
			ServerTime: time.Now().UnixMilli(),
		},
		MutateResults: items,
	}
}

func BuildViewResponse(
	ctx context.Context,
	items []*commonv1.ViewResult,
) *commonv1.ViewResponse {

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	return &commonv1.ViewResponse{
		Metadata: &commonv1.Metadata{
			TraceId:    requestctx.TraceID,
			RequestId:  requestctx.RequestID,
			Locale:     locatectx.Locale,
			Region:     locatectx.Region,
			Degraded:   false,
			ServerTime: time.Now().UnixMilli(),
		},
		ViewResults: items,
	}
}
