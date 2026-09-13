package mapper

import (
	comv1rs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
	dp "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/application/dispatcher"
	"github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

func ToProtoViolations(
	violations []aerrs.Violation,
) []*comv1rs.Violation {
	if len(violations) == 0 {
		return nil
	}

	result := make([]*comv1rs.Violation, 0, len(violations))

	for _, violation := range violations {
		result = append(result, &comv1rs.Violation{
			Field:   violation.Field,
			Message: violation.Message,
			Reason:  violation.Code,
			Hint:    violation.Hint,
		})
	}

	return result
}

func ToAppViolations(
	violations []*comv1rs.Violation,
) []aerrs.Violation {
	if len(violations) == 0 {
		return nil
	}

	result := make([]aerrs.Violation, 0, len(violations))

	for _, violation := range violations {
		if violation == nil {
			continue
		}

		result = append(result, aerrs.Violation{
			Field:   violation.Field,
			Code:    violation.Reason,
			Message: violation.Message,
			Hint:    violation.Hint,
		})
	}

	return result
}

func ToAppErrorInfo(
	opID string,
	appErr *apperrors.AppError,
) *comv1rs.AppErrorInfo {
	if appErr == nil {
		return nil
	}

	return &comv1rs.AppErrorInfo{
		Metadata: &comv1rs.MetadataReponse{
			OpId: opID,
		},
		Reason:     appErr.Code,
		Domain:     appErr.Domain,
		Layer:      appErr.Layer,
		Message:    appErr.Message,
		Violations: ToProtoViolations(appErr.Violations),
	}
}

func ToDispatchSummary(
	multipleErr *dp.MultipleError,
) *comv1rs.DispatchSummary {
	if multipleErr == nil {
		return nil
	}

	total := multipleErr.Total()

	return &comv1rs.DispatchSummary{
		Total:     int32(total),
		Succeeded: 0,
		Failed:    int32(total),
	}
}
