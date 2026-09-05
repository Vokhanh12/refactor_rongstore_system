package mapper

import (
	"context"
	"time"

	protos "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	dp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/dispatcher"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/protobuf/types/known/anypb"
)

func BuildMutateResult(ctx context.Context, Operation, err aerrs.AppError) *protos.MutateResult
{


	return &protos.MutateResult{
		OpId:      opID,
		ResourceId: resourceId,
		Data:       &anypb.Any{},
		Success:    false,
		ClientErr:  &protos.ClientError{},
		ServerErr:  &protos.ServerError{},
	}
}


func BuildBaseResponse(ctx context.Context, result *anypb.Any) *protos.BaseResponse {

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	return &protos.BaseResponse{
		Metadata: &protos.Metadata{
			TraceId:    requestctx.TraceID,
			RequestId:  requestctx.RequestID,
			Locale:     locatectx.Locale,
			Region:     locatectx.Region,
			Degraded:   false,
			ServerTime: time.Now().UnixMilli(),
		},
		Data: result,

		Error: &protos.Error{
			Client: &protos.ClientError{
				Code:       "",
				Message:    "",
				Violations: []*protos.Violation{},
			},
		},
	}
}

func BuildDevBaseResponse(ctx context.Context, result *anypb.Any) *protos.BaseResponse {

	requestctx := ctxutil.MustRequest(ctx)
	locatectx := ctxutil.MustLocale(ctx)

	return &protos.BaseResponse{
		Metadata: &protos.Metadata{
			TraceId:    requestctx.TraceID,
			RequestId:  requestctx.RequestID,
			Locale:     locatectx.Locale,
			Region:     locatectx.Region,
			Degraded:   false,
			ServerTime: time.Now().UnixMilli(),
		},
		Data: result,

		Error: &protos.Error{
			Client: &protos.ClientError{
				Code:       "",
				Message:    "",
				Violations: []*protos.Violation{},
			},
			Server: &protos.ServerError{
				Key:          "",
				Severity:     "",
				Retryable:    false,
				Source:       "",
				GrpcCode:     "",
				ClientAction: "",
				ServerAction: "",
			},
		},
	}
}


func BuildMutateResponse(ctx context.Context, results dp.Result, mapActionData func(data any) *anypb.Any) *protos.MutateResponse {

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
		MutateResults: dispatcherResultToProto(results, mapActionData),
	}
}

func BuildBatch[T any, R any](
	items []T,
	decode func(T) (
		dp.Operation,
		*aerrs.AppError,
	),
) (
	[]dp.Operation,
	*aerrs.AppError,
) {

	out := make(
		[]dp.Operation,
		0,
		len(items),
	)

	for _, item := range items {

		operation, err := decode(item)
		if err != nil {
			return nil, err
		}

		out = append(
			out,
			operation,
		)
	}

	return out, nil
}
