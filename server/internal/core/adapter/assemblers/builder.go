package assemblers

import (
	"context"
	"time"

	protos "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	dp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/dispatcher"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
	"google.golang.org/protobuf/types/known/anypb"
)

func BuildResponse(ctx context.Context, results dp.Result, mapActionData func(data any) *anypb.Any) *protos.MutateResponse {

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

func BuildBatch[T any](
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
