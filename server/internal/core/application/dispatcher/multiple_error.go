package dispatcher

import (
	"errors"

	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type MultipleError struct {
	Errors []*OperationError
}

func NewMultipleError(errs ...error) *MultipleError {
	m := &MultipleError{}

	for _, err := range errs {
		m.Append(err)
	}

	if m.Empty() {
		return nil
	}

	return m
}

func (e *MultipleError) Append(err error) {
	if e == nil || err == nil {
		return
	}

	e.Errors = append(e.Errors, err)
}

func (e *MultipleError) Empty() bool {
	return e == nil || len(e.Errors) == 0
}

func (e *MultipleError) Error() string {
	return "multiple errors"
}

func (e *MultipleError) Unwrap() []error {
	if e == nil {
		return nil
	}

	return e.Errors
}

// ============================================================
// Error information
// ============================================================

func (e *MultipleError) FirstError() error {
	if e == nil || len(e.Errors) == 0 {
		return nil
	}

	return e.Errors[0]
}

func (e *MultipleError) FirstAppError() *aerr.AppError {
	if e == nil {
		return nil
	}

	for _, err := range e.Errors {
		var appErr *aerr.AppError

		if errors.As(err, &appErr) {
			return appErr
		}
	}

	return nil
}

func (e *MultipleError) AppErrors() []*aerr.AppError {
	if e == nil {
		return nil
	}

	appErrors := make(
		[]*aerr.AppError,
		0,
		len(e.Errors),
	)

	for _, err := range e.Errors {
		var appErr *aerr.AppError

		if !errors.As(err, &appErr) {
			continue
		}

		appErrors = append(
			appErrors,
			appErr,
		)
	}

	return appErrors
}

// ============================================================
// Statistics
// ============================================================

func (e *MultipleError) Total() int {
	if e == nil {
		return 0
	}

	return len(e.Errors)
}
