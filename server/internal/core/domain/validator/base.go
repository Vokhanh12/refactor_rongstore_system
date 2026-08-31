package validator

import (
	"net/mail"
	"regexp"
	"strings"

	domain "github.com/vokhanh12/refactor-rongstore-system/server/internal/core/errors"
	err "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type Validator struct {
	violations []err.Violation
}

func New() *Validator {
	return &Validator{
		violations: make([]err.Violation, 0),
	}
}

// ============================
// INTERNAL
// ============================

func (v *Validator) add(reason err.Violation, field string) {
	v.violations = append(
		v.violations,
		err.NewDetail(reason, err.WithField(field)),
	)
}

// ============================
// REQUIRED / NULL
// ============================

func (v *Validator) Required(field string, value string) *Validator {
	if strings.TrimSpace(value) == "" {
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

func (v *Validator) OptionalString(field string, value *string) *Validator {
	if value != nil && strings.TrimSpace(*value) == "" {
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
// EMAIL
// ============================

func (v *Validator) Email(field string, value string) *Validator {
	value = strings.TrimSpace(value)

	_, err := mail.ParseAddress(value)
	if err != nil {
		v.add(domain.REASON_VAL_INVALID_FORMAT, field)
	}

	return v
}

// ============================
// PHONE
// ============================

var vietnamesePhonePattern = regexp.MustCompile(
	`^(?:\+84|0)(?:3|5|7|8|9)[0-9]{8}$`,
)

func (v *Validator) Phone(field string, value string) *Validator {
	value = strings.TrimSpace(value)

	if !vietnamesePhonePattern.MatchString(value) {
		v.add(domain.REASON_VAL_INVALID_FORMAT, field)
	}

	return v
}

// ============================
// PHONE OR EMAIL
// ============================

func (v *Validator) EmailOrPhone(field string, value string) *Validator {
	value = strings.TrimSpace(value)

	_, emailErr := mail.ParseAddress(value)
	isEmail := emailErr == nil

	isPhone := vietnamesePhonePattern.MatchString(value)

	if !isEmail && !isPhone {
		v.add(domain.REASON_VAL_INVALID_FORMAT, field)
	}

	return v
}

// ============================
// RESULT
// ============================

func (v *Validator) Err() error {
	if len(v.violations) == 0 {
		return nil
	}

	return err.New(
		domain.VALIDATION_FAILED,
		err.WithAppendViolations(v.violations),
	)
}
