package assemblers

import (
	"context"
	"time"

	protos "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	uc "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/usecase"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/protobuf/types/known/anypb"
)

func BuildMutateResult(opID string, data any, err *aerrs.AppError) dtos.MutateResultItemDTO {

	return dtos.MutateResultItemDTO{
		OpID:  opID,
		Data:  data,
		Error: appErrorToDTO(err),
	}
}

func BuildViewResult(opID string, data any, err *aerrs.AppError) dtos.ViewResultItemDTO {

	return dtos.ViewResultItemDTO{
		OpID:  opID,
		Items: data,
		Error: appErrorToDTO(err),
	}
}

func BuildMutateResponse(ctx context.Context, results dtos.MutateResultDTO, mapActionData func(data any) *anypb.Any) *protos.MutateResponse {

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
		MutateResults: mutateResultToProto(results, mapActionData),
	}
}

func BuildBatch[T any, R any](
	items []T,
	decode func(T) (R, *aerrs.AppError),
	getID func(T) string,
) []uc.Operation[R] {

	out := make([]uc.Operation[R], 0, len(items))

	for _, item := range items {

		payload, err := decode(item)

		out = append(out, uc.Operation[R]{
			OpID:    getID(item),
			Payload: payload,
			Error:   err,
		})
	}

	return out
}

// func BuildViewResponse(
// 	ctx context.Context,
// 	items []*protos.ViewResult,
// ) *protos.ViewResponse {

// 	requestctx := ctxutil.MustRequest(ctx)
// 	locatectx := ctxutil.MustLocale(ctx)

// 	return &protos.ViewResponse{
// 		Metadata: &protos.Metadata{
// 			TraceId:    requestctx.TraceID,
// 			RequestId:  requestctx.RequestID,
// 			Locale:     locatectx.Locale,
// 			Region:     locatectx.Region,
// 			Degraded:   false,
// 			ServerTime: time.Now().UnixMilli(),
// 		},
// 		ViewResults: items,
// 	}
// }
