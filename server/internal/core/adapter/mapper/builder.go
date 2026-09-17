package mapper

import (
	"context"
	"time"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	comv1rs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	dp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/dispatcher"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/ctxutil"
)

func BuildSuccessResponse(
	ctx context.Context,
	data proto.Message,
) (*comv1rs.SuccessResponse, error) {

	requestctx := ctxutil.MustRequest(ctx)

	anyData, err := anypb.New(data)
	if err != nil {
		return nil, err
	}

	return &comv1rs.SuccessResponse{
		Metadata: &comv1rs.ResponseMetadata{
			TraceId:    requestctx.TraceID,
			RequestId:  requestctx.RequestID,
			Degraded:   false,
			ServerTime: time.Now().UnixMilli(),
		},
		Data: anyData,
	}, nil
}

func BuildMutateResponse(
	ctx context.Context,
	results []*comv1rs.MutateResult,
	multiErr *dp.MultipleError,
) (*comv1rs.MutateResponse, error) {

	resp := &comv1rs.MutateResponse{
		MutateResults: results,
	}

	if multiErr != nil && !multiErr.Empty() {
		return resp, multiErr
	}

	return resp, nil
}

func BuildViewResponse(
	ctx context.Context,
	results []*comv1rs.ViewResult,
	multiErr *dp.MultipleError,
) (*comv1rs.ViewResponse, error) {

	resp := &comv1rs.ViewResponse{
		ViewResults: results,
	}

	if multiErr != nil && !multiErr.Empty() {
		return resp, multiErr
	}

	return resp, nil
}
