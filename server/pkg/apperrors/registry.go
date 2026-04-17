package apperrors

// New creates a new AppError from a template with optional modifiers
func New(template AppError, opts ...func(*AppError)) *AppError {
	e := copyError(template)
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// NewDetail creates a new AppErrorDetail from a template
func NewDetail(template AppErrorDetail, opts ...func(*AppErrorDetail)) *AppErrorDetail {
	e := copyDetail(template)
	for _, opt := range opts {
		opt(e)
	}
	return e
}

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

func WithField(value string) func(*AppErrorDetail) {
	return func(e *AppErrorDetail) {
		e.Field = value
	}
}

func WithMessageDetail(msg string) func(*AppErrorDetail) {
	return func(e *AppErrorDetail) {
		e.Message = msg
	}
}

func WithData(data map[string]interface{}) func(*AppError) {
	return func(e *AppError) {
		e.Data = copyDataMap(data)
	}
}

func WithCauseDetail(err error) func(*AppError) {
	return func(e *AppError) {
		e.causeDetail = err
	}
}

func WithErrorDetails(errors []AppErrorDetail) func(*AppError) {
	return func(e *AppError) {
		e.errorDetails = copyDetails(errors)
	}
}

func WithAppendErrorDetails(details []AppErrorDetail) func(*AppError) {
	return func(e *AppError) {
		e.errorDetails = append(e.errorDetails, details...)
	}
}

func WithAppendErrorDetail(detail AppErrorDetail) func(*AppError) {
	return func(e *AppError) {
		e.errorDetails = append(e.errorDetails, detail)
	}
}

// ---- internal helpers ----

func copyError(src AppError) *AppError {
	dst := src
	dst.Data = copyDataMap(src.Data)
	if src.Tags != nil {
		dst.Tags = append([]string(nil), src.Tags...)
	}
	dst.errorDetails = copyDetails(src.errorDetails)
	return &dst
}

func copyDetail(src AppErrorDetail) *AppErrorDetail {
	dst := src
	return &dst
}

func copyDataMap(src map[string]interface{}) map[string]interface{} {
	if src == nil {
		return nil
	}

	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func copyDetails(src []AppErrorDetail) []AppErrorDetail {
	if src == nil {
		return nil
	}

	dst := make([]AppErrorDetail, len(src))
	copy(dst, src)
	return dst
}
