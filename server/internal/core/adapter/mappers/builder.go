package mappers

import (
	"context"
	"time"

	protos "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func BuildMutateResult(opID string, data any, err aerrs.AppError) dtos.MutateResultItemDTO {
	return dtos.MutateResultItemDTO{
		OpID:  opID,
		Data:  data,
		Error: AppErrorToDTO(err),
	}
}

func BuildTransportErrorResponse(ctx context.Context, err *aerrs.AppError) (*protos.BaseResponse, error) {

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	return &protos.BaseResponse{
		Success: false,
		Metadata: &protos.Metadata{
			TraceId:    requestctx.TraceID,
			RequestId:  requestctx.RequestID,
			Locale:     locatectx.Locale,
			Region:     locatectx.Region,
			Degraded:   false,
			ServerTime: time.Now().UnixMilli(),
		},
		Error: AppErrorToDTO(err),

		// &protos.Error{
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

func BuildMutateErrorResponse(ctx context.Context, errs []aerrs.AppError) *protos.MutateResponse {

	results := make([]*protos.MutateResult, 0, len(errs))

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	for _, err := range errs {
		results = append(results, &protos.MutateResult{
			Success: false,
			Error: &protos.Error{
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

	return &protos.MutateResponse{
		Metadata: &protos.Metadata{
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
func BuildSuccessResponse(ctx context.Context, data proto.Message) (*protos.BaseResponse, error) {
	anyData, err := anypb.New(data)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to marshal response")
	}

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	return &protos.BaseResponse{
		Success: true,
		Data:    anyData,
		Metadata: &protos.Metadata{
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
// func FromError(ctx context.Context, err error) *protos.BaseResponse {
// 	be, _ := aerrs.GetBusinessError(err)
// 	return BuildErrorResponse(ctx, be)
// }

func BuildMutateResponse(ctx context.Context, results dtos.MutateResultDTO) *protos.MutateResponse {

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	return &protos.MutateResponse{
		Metadata: &protos.Metadata{
			TraceId:    requestctx.TraceID,
			RequestId:  requestctx.RequestID,
			Locale:     locatectx.Locale,
			Region:     locatectx.Region,
			Degraded:   false,
			ServerTime: time.Now().UnixMilli(),
		},
		MutateResults: MutateResultToProto(results),
	}
}

func BuildViewResponse(
	ctx context.Context,
	items []*protos.ViewResult,
) *protos.ViewResponse {

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	return &protos.ViewResponse{
		Metadata: &protos.Metadata{
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
