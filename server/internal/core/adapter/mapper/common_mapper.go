package mapper

import (
	commonv1 "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/common/v1"
	pkgcommonv1 "github.com/vokhanh12/refactor-rongstore-system/server/pkg/common/v1"
)

func ErrorToProto(dto pkgcommonv1.AppError) *commonv1.Error {

	details := make([]*commonv1.ErrorDetail, 0, len(dto.Details))

	for _, d := range dto.Details {
		details = append(details, ErrorDetailToProto(d))
	}

	return &commonv1.Error{
		Code:         dto.Code,
		Key:          dto.Key,
		Message:      dto.Message,
		Severity:     dto.Severity,
		Retryable:    dto.Retryable,
		Source:       dto.Source,
		GrpcCode:     dto.GRPCCode,
		HttpStatus:   dto.HTTPStatus,
		ClientAction: dto.ClientAction,
		ServerAction: dto.ServerAction,
		Details:      details,
	}
}

func ErrorDetailToProto(dto pkgcommonv1.AppErrorDetail) *commonv1.ErrorDetail {
	return &commonv1.ErrorDetail{
		Field:   dto.Field,
		Message: dto.Message,
		Code:    dto.Code,
		Hint:    dto.Hint,
	}
}

func MutateResultToProto(dto pkgcommonv1.MutateResult) *commonv1.MutateResult {

	items := make([]*commonv1.MutateResultItem, 0, len(dto.Items))

	for _, it := range dto.Items {
		items = append(items, MutateResultItemToProto(it))
	}

	return &commonv1.MutateResult{
		Items: items,
	}
}

func MutateResultItemToProto(dto pkgcommonv1.MutateResultItem) *commonv1.MutateResultItem {

	return &commonv1.MutateResultItem{
		OpId:    dto.OpID,
		Data:    dto.Data,
		Success: dto.Success,
		Error:   ErrorToProto(*dto.Error),
	}
}
