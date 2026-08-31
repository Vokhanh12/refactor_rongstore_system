package apperrors

func WithMessage(msg string) func(*AppError) {
	return func(e *AppError) {
		e.Message = msg
	}
}

func WithSource(value string) func(*AppError) {
	return func(e *AppError) {
		e.Source = value
	}
}

func WithData(data map[string]interface{}) func(*AppError) {
	return func(e *AppError) {
		e.Data = copyDataMap(data)
	}
}

func WithErr(err error) func(*AppError) {
	return func(e *AppError) {
		e.Err = err
	}
}

func WithViolation(details []Violation) func(*AppError) {
	return func(e *AppError) {
		e.Violations = copyDetails(details)
	}
}

func WithAppendViolations(details []Violation) func(*AppError) {
	return func(e *AppError) {
		e.Violations = append(e.Violations, details...)
	}
}

func WithAppendViolation(detail Violation) func(*AppError) {
	return func(e *AppError) {
		e.Violations = append(e.Violations, detail)
	}
}
