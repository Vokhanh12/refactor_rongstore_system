package mappers

import (
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
	"google.golang.org/protobuf/types/known/anypb"

	protos "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

// ============================================================
// TO → PROTO
// ============================================================

func appErrorToProto(it *dtos.ErrorDTO) *protos.Error {

	if it == nil {
		return nil
	}

	items := make([]*protos.ErrorDetail, 0, len(it.External.Details))

	for _, d := range it.External.Details {
		items = append(items, appErrorDetailToProto(d))
	}

	return &protos.Error{
		External: &protos.ExternalError{
			Code:    it.External.Code,
			Message: it.External.Message,
			Details: items,
		},
		Internal: &protos.InternalError{
			Key:          it.Internal.Key,
			Severity:     it.Internal.Severity,
			Retryable:    it.Internal.Retryable,
			Source:       it.Internal.Source,
			GrpcCode:     it.Internal.GRPCCode,
			ClientAction: it.Internal.ClientAction,
			ServerAction: it.Internal.ServerAction,
		},
	}
}

func appErrorDetailToProto(it dtos.ErrorDetailDTO) *protos.ErrorDetail {
	return &protos.ErrorDetail{
		Field:   it.Field,
		Message: it.Message,
		Code:    it.Code,
		Hint:    it.Hint,
	}
}

func mutateResultToProto(dto dtos.MutateResultDTO, mapActionData func(data any) *anypb.Any) *protos.MutateResult {

	items := make([]*protos.MutateResultItem, 0, len(dto.Items))

	for _, it := range dto.Items {
		items = append(items, &protos.MutateResultItem{
			OpId:  it.OpID,
			Data:  mapActionData(it.Data),
			Error: appErrorToProto(it.Error),
		})
	}

	return &protos.MutateResult{
		Items: items,
	}
}

// ============================================================
// TO → DTO
// ============================================================

func appErrorToDTO(it *aerrs.AppError) *dtos.ErrorDTO {
	if it == nil {
		return nil
	}

	details := make([]dtos.ErrorDetailDTO, 0)
	if it.ErrorDetails != nil {
		for _, d := range it.ErrorDetails {
			details = append(details, appErrorDetailToDTO(d))
		}
	}

	return &dtos.ErrorDTO{
		External: dtos.ExternalErrorDTO{
			Code:    it.Code,
			Message: it.Message,
			Details: details,
		},
		Internal: dtos.InternalErrorDTO{
			Code:         it.Code,
			Key:          it.Key,
			Message:      it.Message,
			Severity:     it.Severity,
			Retryable:    it.Retryable,
			Source:       it.Source,
			Component:    it.Component,
			GRPCCode:     it.GRPCCode,
			ClientAction: it.ClientAction,
			ServerAction: it.ServerAction,
		},
	}
}

func appErrorDetailToDTO(it aerrs.AppErrorDetail) dtos.ErrorDetailDTO {
	return dtos.ErrorDetailDTO{
		Field:   it.Field,
		Message: it.Message,
		Code:    it.Code,
		Hint:    it.Hint,
	}
}
