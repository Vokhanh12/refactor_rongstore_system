package mappers

import (
	protos "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

// ============================================================
// COMMAND / QUERY / DTO → PROTO
// ============================================================

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
		Error:   ErrorToProto(*dto.Error),
	}
}

// ============================================================
// LIBARY → COMMAND / QUERY / DTO
// ============================================================

func ExternalAppErrorToDTO(err *aerrs.AppError) *dtos.ExternalAppErrorDTO {

	if err == nil {
		return nil
	}

	if err.ErrorDetails == nil {
		return &dtos.ExternalAppErrorDTO{
			Code:    err.Code,
			Message: err.Message,
			Details: nil,
		}
	}

	details := make([]dtos.ExternalAppErrorDetailDTO, 0, len(*err.ErrorDetails))

	for _, d := range *err.ErrorDetails {
		details = append(details, externalErrorDetailToDTO(d))
	}

	return &dtos.ExternalAppErrorDTO{
		Code:    err.Code,
		Message: err.Message,
		Details: &details,
	}
}

func externalErrorDetailToDTO(errdetail aerrs.AppErrorDetail) dtos.ExternalAppErrorDetailDTO {
	return dtos.ExternalAppErrorDetailDTO{
		Field:   errdetail.Field,
		Message: errdetail.Message,
		Code:    errdetail.Code,
		Hint:    errdetail.Hint,
	}
}

func InternalAppErrorToDTO(err *aerrs.AppError) *dtos.InternalAppErrorDTO {

	if err == nil {
		return nil
	}

	if err.ErrorDetails == nil {
		return &dtos.InternalAppErrorDTO{
			Code:         err.Code,
			Status:       err.Status,
			GRPCCode:     err.GRPCCode,
			Key:          err.Key,
			Cause:        err.Cause,
			ClientAction: err.ClientAction,
			ServerAction: err.ServerAction,
			Source:       err.Source,
			Component:    err.Component,
			Tags:         err.Tags,
			Message:      err.Message,
			Data:         err.Data,
			Severity:     err.Severity,
			Expected:     err.Expected,
			Retryable:    err.Retryable,
			CauseDetail:  err.CauseDetail,
			ErrorDetails: nil,
		}
	}

	details := make([]dtos.InternalAppErrorDetailDTO, 0, len(*err.ErrorDetails))

	for _, d := range *err.ErrorDetails {
		details = append(details, internalErrorDetailToDTO(d))
	}

	return &dtos.InternalAppErrorDTO{
		Code:         err.Code,
		Status:       err.Status,
		GRPCCode:     err.GRPCCode,
		Key:          err.Key,
		Cause:        err.Cause,
		ClientAction: err.ClientAction,
		ServerAction: err.ServerAction,
		Source:       err.Source,
		Component:    err.Component,
		Tags:         err.Tags,
		Message:      err.Message,
		Data:         err.Data,
		Severity:     err.Severity,
		Expected:     err.Expected,
		Retryable:    err.Retryable,
		CauseDetail:  err.CauseDetail,
		ErrorDetails: &details,
	}
}

func internalErrorDetailToDTO(errdetail aerrs.AppErrorDetail) dtos.InternalAppErrorDetailDTO {
	return dtos.InternalAppErrorDetailDTO{
		Field:   errdetail.Field,
		Message: errdetail.Message,
		Code:    errdetail.Code,
		Hint:    errdetail.Hint,
	}
}

// ============================================================
// LIBARY → PROTO
// ============================================================

func AppErrorToProto(err *aerrs.AppError) *protos.AppError {

	if err == nil {
		return nil
	}

	if err.ErrorDetails == nil {
		return &protos.Error{
			Code:    err.Code,
			Message: err.Message,
			Details: nil,
		}
	}

	details := make([]protos.AppErrorDetail, 0, len(*err.ErrorDetails))

	for _, d := range *err.ErrorDetails {
		details = append(details, protos.AppErrorDetail{
			Field:   d.Field,
			Message: d.Message,
			Code:    d.Code,
			Hint:    d.Hint,
		})
	}

	return &protos.AppError{
		Code:    err.Code,
		Message: err.Message,
		Details: &details,
	}
}
