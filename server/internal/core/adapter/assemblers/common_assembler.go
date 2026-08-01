package assemblers

import (
	"google.golang.org/protobuf/types/known/anypb"

	protos "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	dp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/dispatcher"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

func dispatcherResultToProto(result dp.Result, mapActionData func(data any) *anypb.Any) *protos.MutateResult {

	items := make([]*protos.MutateResultItem, 0, len(result.Items))

	for _, it := range result.Items {
		items = append(items, &protos.MutateResultItem{
			OpId:  it.OpID,
			Data:  mapActionData(it.Data),
			Error: errToProto(it.Error),
		})
	}

	return &protos.MutateResult{
		Items: items,
	}
}

func errToProto(it *aerrs.AppError) *protos.Error {

	if it == nil {
		return nil
	}

	items := make([]*protos.ErrorDetail, 0, len(it.Violations))

	for _, d := range it.Violations {
		items = append(items, violationToProto(d))
	}

	return &protos.Error{
		External: &protos.ExternalError{
			Code:    it.Code,
			Message: it.Message,
			Details: items,
		},
		Internal: &protos.InternalError{
			Key:          it.Key,
			Severity:     it.Severity,
			Retryable:    it.Retryable,
			Source:       it.Source,
			GrpcCode:     it.GRPCCode,
			ClientAction: it.ClientAction,
			ServerAction: it.ServerAction,
		},
	}
}

func violationToProto(it aerrs.Violation) *protos.ErrorDetail {
	return &protos.ErrorDetail{
		Field:   it.Field,
		Message: it.Message,
		Code:    it.Code,
		Hint:    it.Hint,
	}
}
