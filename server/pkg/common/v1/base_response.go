package commonv1

import (
	"context"
	"time"

	"github.com/golang/protobuf/ptypes/any"
	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/apperrors"
	util "github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type BaseResponse struct {
	Success    bool              `json:"success"`
	Data       *any.Any          `json:"data,omitempty"`
	Metadata   *Metadata         `json:"metadata,omitempty"`
	Error      *AppError         `json:"error,omitempty"`
	Pagination *Pagination       `json:"pagination,omitempty"`
	Warnings   []*Warning        `json:"warnings,omitempty"`
	Details    map[string]string `json:"details,omitempty"`
}

type MutateResponse struct {
	Metadata      *Metadata       `json:"metadata,omitempty"`
	MutateResults []*MutateResult `json:"mutate_results,omitempty"`
}

type ViewResponse struct {
	Metadata    *Metadata     `json:"metadata,omitempty"`
	ViewResults []*ViewResult `json:"view_results,omitempty"`
}

func BuildTransportErrorResponse(
	ctx context.Context,
	err *aerrs.AppError,
) (*commonv1.BaseResponse, error) {

	requestctx := util.MustRequest(ctx)
	locatectx := util.MustLocale(ctx)

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
		Error: &commonv1.Error{
			//Code:         err.Code,
			//Key:          err.Key,
			Message: err.Message,
			//Severity:     err.Severity,
			//Retryable:    err.Retryable,
			//Source:       err.Source,
			GrpcCode:   err.GRPCCode,
			HttpStatus: int32(err.Status),
			//ClientAction: err.ClientAction,
			//ServerAction: err.ServerAction,
			Details: aerrs.MapAppErrorDetailsToProto(err.GetErrorDetails()),
		},
	}, nil

}

func BuildMutateErrorResponse(
	ctx context.Context,
	errs []aerrs.AppError,
) *commonv1.MutateResponse {

	results := make([]*commonv1.MutateResult, 0, len(errs))

	requestctx := util.MustRequest(ctx)
	locatectx := util.MustLocale(ctx)

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

	requestctx := util.MustRequest(ctx)
	locatectx := util.MustLocale(ctx)

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

	requestctx := util.MustRequest(ctx)
	locatectx := util.MustLocale(ctx)

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

	requestctx := util.MustRequest(ctx)
	locatectx := util.MustLocale(ctx)

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
