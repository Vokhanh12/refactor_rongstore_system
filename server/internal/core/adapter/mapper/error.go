package mapper

import (
	comv1rs "github.com/vokhanh12/refactor-rongstore-system/server/gen/proto/core/common/v1/resources"
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
