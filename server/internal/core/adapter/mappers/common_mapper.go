package mappers

import (
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"

	protos "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

// ============================================================
// TO → PROTO
// ============================================================

func AppErrorToProto(it dtos.ErrorDTO) *protos.Error {

	items := make([]*protos.ErrorDetail, 0, len(it.External.Details))

	for _, d := range it.External.Details {
		items = append(items, AppErrorDetailToProto(d))
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

func AppErrorDetailToProto(it dtos.ErrorDetailDTO) *protos.ErrorDetail {
	return &protos.ErrorDetail{
		Field:   it.Field,
		Message: it.Message,
		Code:    it.Code,
		Hint:    it.Hint,
	}
}

func MutateResultToProto(dto dtos.MutateResultDTO) *protos.MutateResult {

	items := make([]*protos.MutateResultItem, 0, len(dto.Items))

	for _, it := range dto.Items {
		items = append(items, MutateResultItemToProto(it))
	}

	return &protos.MutateResult{
		Items: items,
	}
}

func MutateResultItemToProto(dto dtos.MutateResultItemDTO) *protos.MutateResultItem {
	return &protos.MutateResultItem{
		OpId:    dto.OpID,
		Data:    dto.Data,
		Success: dto.Success,
		Error:   AppErrorToProto(dto.Error),
	}
}

// ============================================================
// TO → DTO
// ============================================================

func AppErrorToDTO(it *aerrs.AppError) *dtos.ErrorDTO {

	if it == nil {
		return nil
	}

	var items []dtos.ErrorDetailDTO

	if it.ErrorDetails != nil {
		items = make([]dtos.ErrorDetailDTO, 0, len(*it.ErrorDetails))

		for _, d := range *it.ErrorDetails {
			items = append(items, AppErrorDetailToDTO(d))
		}
	} else {
		items = []dtos.ErrorDetailDTO{}
	}

	return &dtos.ErrorDTO{
		External: dtos.ExternalErrorDTO{
			Code:    it.Code,
			Message: it.Message,
			Details: items,
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

func AppErrorDetailToDTO(it aerrs.AppErrorDetail) dtos.ErrorDetailDTO {
	return dtos.ErrorDetailDTO{
		Field:   it.Field,
		Message: it.Message,
		Code:    it.Code,
		Hint:    it.Hint,
	}
}
