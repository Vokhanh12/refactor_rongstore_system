package dispatcher

import (
	"errors"

	aerr "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"
)

type MultipleError struct {
	Errors    []error
	AppErrors []*aerr.AppError
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
	if err == nil {
		return
	}

	var appErr *aerr.AppError

	if errors.As(err, &appErr) {
		e.AppErrors = append(e.AppErrors, appErr)
		return
	}

	e.Errors = append(e.Errors, err)
}

func (e *MultipleError) Empty() bool {
	return len(e.Errors) == 0 &&
		len(e.AppErrors) == 0
}

func (e *MultipleError) Error() string {
	return "multiple errors"
}

func (e *MultipleError) Unwrap() []error {
	errs := make([]error, 0, len(e.Errors)+len(e.AppErrors))

	errs = append(errs, e.Errors...)

	for _, err := range e.AppErrors {
		errs = append(errs, err)
	}

	return errs
}
