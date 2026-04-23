package validator

import (
	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type Validator struct {
	details []aerrs.AppErrorDetail
}

func New() *Validator {
	return &Validator{
		details: make([]aerrs.AppErrorDetail, 0),
	}
}

// ============================
// INTERNAL
// ============================

func (v *Validator) add(reason aerrs.AppErrorDetail, field string) {
	v.details = append(v.details,
		aerrs.NewDetail(reason, aerrs.WithField(field)),
	)
}

// ============================
// REQUIRED / NULL
// ============================

func (v *Validator) Required(field string, value string) *Validator {
	if value == "" {
		v.add(domain.REASON_VAL_REQUIRED, field)
	}
	return v
}

func (v *Validator) NotNil(field string, value any) *Validator {
	if value == nil {
		v.add(domain.REASON_VAL_NULL, field)
	}
	return v
}

// Optional nhưng nếu có thì phải valid
func (v *Validator) OptionalString(field string, value *string) *Validator {
	if value != nil && *value == "" {
		v.add(domain.REASON_VAL_REQUIRED, field)
	}
	return v
}

// ============================
// FORMAT / TYPE
// ============================

func (v *Validator) Format(field string, valid bool) *Validator {
	if !valid {
		v.add(domain.REASON_VAL_INVALID_FORMAT, field)
	}
	return v
}

func (v *Validator) Type(field string, valid bool) *Validator {
	if !valid {
		v.add(domain.REASON_VAL_INVALID_TYPE, field)
	}
	return v
}

func (v *Validator) Enum(field string, valid bool) *Validator {
	if !valid {
		v.add(domain.REASON_VAL_INVALID_ENUM, field)
	}
	return v
}

// ============================
// RANGE / SIZE
// ============================

func (v *Validator) MinInt(field string, value int, min int) *Validator {
	if value < min {
		v.add(domain.REASON_VAL_MIN, field)
	}
	return v
}

func (v *Validator) MaxInt(field string, value int, max int) *Validator {
	if value > max {
		v.add(domain.REASON_VAL_MAX, field)
	}
	return v
}

func (v *Validator) RangeInt(field string, value int, min int, max int) *Validator {
	if value < min || value > max {
		v.add(domain.REASON_VAL_OUT_OF_RANGE, field)
	}
	return v
}

func (v *Validator) Uint8Max(field string, value uint8, max uint8) *Validator {
	if value > max {
		v.add(domain.REASON_VAL_OUT_OF_RANGE, field)
	}
	return v
}

// ============================
// STRING / LENGTH
// ============================

func (v *Validator) MinLen(field string, value string, min int) *Validator {
	if len(value) < min {
		v.add(domain.REASON_VAL_TOO_SHORT, field)
	}
	return v
}

func (v *Validator) MaxLen(field string, value string, max int) *Validator {
	if len(value) > max {
		v.add(domain.REASON_VAL_TOO_LONG, field)
	}
	return v
}

func (v *Validator) Pattern(field string, valid bool) *Validator {
	if !valid {
		v.add(domain.REASON_VAL_INVALID_PATTERN, field)
	}
	return v
}

// ============================
// BUSINESS
// ============================

// Application layer thường dùng
func (v *Validator) Duplicate(field string, isDuplicate bool) *Validator {
	if isDuplicate {
		v.add(domain.REASON_VAL_DUPLICATE, field)
	}
	return v
}

func (v *Validator) Conflict(field string, conflict bool) *Validator {
	if conflict {
		v.add(domain.REASON_VAL_CONFLICT, field)
	}
	return v
}

func (v *Validator) NotAllowed(field string, notAllowed bool) *Validator {
	if notAllowed {
		v.add(domain.REASON_VAL_NOT_ALLOWED, field)
	}
	return v
}

func (v *Validator) ForbiddenState(field string, invalid bool) *Validator {
	if invalid {
		v.add(domain.REASON_VAL_FORBIDDEN_STATE, field)
	}
	return v
}

// ============================
// EXISTENCE
// ============================

func (v *Validator) NotFound(field string, notFound bool) *Validator {
	if notFound {
		v.add(domain.REASON_VAL_NOT_FOUND, field)
	}
	return v
}

func (v *Validator) AlreadyExists(field string, exists bool) *Validator {
	if exists {
		v.add(domain.REASON_VAL_ALREADY_EXISTS, field)
	}
	return v
}

// ============================
// RELATION
// ============================

func (v *Validator) Reference(field string, valid bool) *Validator {
	if !valid {
		v.add(domain.REASON_VAL_INVALID_REFERENCE, field)
	}
	return v
}

func (v *Validator) Dependency(field string, missing bool) *Validator {
	if missing {
		v.add(domain.REASON_VAL_DEPENDENCY_MISSING, field)
	}
	return v
}

// ============================
// TIME
// ============================

func (v *Validator) TimeRange(field string, valid bool) *Validator {
	if !valid {
		v.add(domain.REASON_VAL_INVALID_TIME_RANGE, field)
	}
	return v
}

func (v *Validator) Expired(field string, expired bool) *Validator {
	if expired {
		v.add(domain.REASON_VAL_EXPIRED, field)
	}
	return v
}

// ============================
// SECURITY
// ============================

func (v *Validator) Unsafe(field string, unsafe bool) *Validator {
	if unsafe {
		v.add(domain.REASON_VAL_UNSAFE_INPUT, field)
	}
	return v
}

// ============================
// GENERIC
// ============================

func (v *Validator) Invalid(field string, invalid bool) *Validator {
	if invalid {
		v.add(domain.REASON_VAL_INVALID, field)
	}
	return v
}

// ============================
// RESULT
// ============================

func (v *Validator) Err() *aerrs.AppError {
	if len(v.details) == 0 {
		return nil
	}

	return aerrs.New(
		domain.VALIDATION_FAILED,
		aerrs.WithAppendErrorDetails(v.details),
	)
}
