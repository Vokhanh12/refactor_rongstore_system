package mapper

import (
	v "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

var validationCodeMap = map[string]aerr.Violation{
	"required":      v.REASON_VAL_REQUIRED,
	"required_with": v.REASON_VAL_REQUIRED,

	"string.email":        v.REASON_VAL_INVALID_FORMAT,
	"string.uri":          v.REASON_VAL_INVALID_FORMAT,
	"string.uuid":         v.REASON_VAL_INVALID_FORMAT,
	"string.hostname":     v.REASON_VAL_INVALID_FORMAT,
	"string.ip":           v.REASON_VAL_INVALID_FORMAT,
	"string.ipv4":         v.REASON_VAL_INVALID_FORMAT,
	"string.ipv6":         v.REASON_VAL_INVALID_FORMAT,
	"string.pattern":      v.REASON_VAL_INVALID_FORMAT,
	"string.prefix":       v.REASON_VAL_INVALID_FORMAT,
	"string.suffix":       v.REASON_VAL_INVALID_FORMAT,
	"string.contains":     v.REASON_VAL_INVALID_FORMAT,
	"string.not_contains": v.REASON_VAL_INVALID_FORMAT,

	"string.min_len": v.REASON_VAL_TOO_SHORT,
	"string.max_len": v.REASON_VAL_TOO_LONG,

	"int32.gte": v.REASON_VAL_OUT_OF_RANGE,
	"int32.gt":  v.REASON_VAL_OUT_OF_RANGE,
	"int32.lte": v.REASON_VAL_OUT_OF_RANGE,
	"int32.lt":  v.REASON_VAL_OUT_OF_RANGE,

	"int64.gte": v.REASON_VAL_OUT_OF_RANGE,
	"int64.gt":  v.REASON_VAL_OUT_OF_RANGE,
	"int64.lte": v.REASON_VAL_OUT_OF_RANGE,
	"int64.lt":  v.REASON_VAL_OUT_OF_RANGE,

	"uint32.gte": v.REASON_VAL_OUT_OF_RANGE,
	"uint32.gt":  v.REASON_VAL_OUT_OF_RANGE,
	"uint32.lte": v.REASON_VAL_OUT_OF_RANGE,
	"uint32.lt":  v.REASON_VAL_OUT_OF_RANGE,

	"uint64.gte": v.REASON_VAL_OUT_OF_RANGE,
	"uint64.gt":  v.REASON_VAL_OUT_OF_RANGE,
	"uint64.lte": v.REASON_VAL_OUT_OF_RANGE,
	"uint64.lt":  v.REASON_VAL_OUT_OF_RANGE,

	"double.gte": v.REASON_VAL_OUT_OF_RANGE,
	"double.gt":  v.REASON_VAL_OUT_OF_RANGE,
	"double.lte": v.REASON_VAL_OUT_OF_RANGE,
	"double.lt":  v.REASON_VAL_OUT_OF_RANGE,

	"enum.defined_only": v.REASON_VAL_INVALID_ENUM,

	"repeated.min_items": v.REASON_VAL_TOO_SHORT,
	"repeated.max_items": v.REASON_VAL_TOO_LONG,

	"bytes.min_len": v.REASON_VAL_TOO_SHORT,
	"bytes.max_len": v.REASON_VAL_TOO_LONG,

	"message.required": v.REASON_VAL_REQUIRED,
}

func ToValidationViolation(code string) aerr.Violation {
	if violation, ok := validationCodeMap[code]; ok {
		return violation
	}

	return v.REASON_VAL_INVALID_FORMAT
}
