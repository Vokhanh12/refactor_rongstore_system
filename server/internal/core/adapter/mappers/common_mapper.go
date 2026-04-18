package mappers

import (
	protos "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	dtos "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

// ============================================================
// COMMAND / QUERY / DTO → PROTO
// ============================================================

func MutateResultToProto(dto dtos.MutateResult) *protos.MutateResult {

	items := make([]*protos.MutateResultItem, 0, len(dto.Items))

	for _, it := range dto.Items {
		items = append(items, MutateResultItemToProto(it))
	}

	return &protos.MutateResult{
		Items: items,
	}
}

func MutateResultItemToProto(dto dtos.MutateResultItem) *protos.MutateResultItem {

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

func AppErrorToDTO(err *aerrs.AppError) *dtos.AppError {

	if err == nil {
		return nil
	}

	if err.ErrorDetails == nil {
		return &dtos.AppError{
			Code:    err.Code,
			Message: err.Message,
			Details: nil,
		}
	}

	details := make([]dtos.AppErrorDetail, 0, len(*err.ErrorDetails))

	for _, d := range *err.ErrorDetails {
		details = append(details, dtos.AppErrorDetail{
			Field:   d.Field,
			Message: d.Message,
			Code:    d.Code,
			Hint:    d.Hint,
		})
	}

	return &dtos.AppError{
		Code:    err.Code,
		Message: err.Message,
		Details: &details,
	}
}

func ErrorDetailToDTO(errdetail aerrs.AppErrorDetail) *dtos.AppErrorDetail {
	return &dtos.AppErrorDetail{
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
