package apperrors

// ============================
// MAIN ERROR BUILDER
// ============================

func New(template AppError, opts ...func(*AppError)) *AppError {
	e := copyError(template)
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// ============================
// DETAIL BUILDER
// ============================

func NewDetail(template AppErrorDetail, opts ...func(*AppErrorDetail)) AppErrorDetail {
	e := template
	for _, opt := range opts {
		opt(&e)
	}
	return e
}
