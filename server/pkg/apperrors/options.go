package apperrors

// ===== AppError options =====

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

func WithCauseDetail(err error) func(*AppError) {
	return func(e *AppError) {
		e.CauseDetail = err
	}
}

func WithErrorDetails(details []AppErrorDetail) func(*AppError) {
	return func(e *AppError) {
		e.ErrorDetails = copyDetails(details)
	}
}

func WithAppendErrorDetails(details []AppErrorDetail) func(*AppError) {
	return func(e *AppError) {
		e.ErrorDetails = append(e.ErrorDetails, details...)
	}
}

func WithAppendErrorDetail(detail AppErrorDetail) func(*AppError) {
	return func(e *AppError) {
		e.ErrorDetails = append(e.ErrorDetails, detail)
	}
}
