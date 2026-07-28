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
// DETAIL BUILDERP
// ============================

func NewDetail(template Violation, opts ...func(*Violation)) Violation {
	e := template
	for _, opt := range opts {
		opt(&e)
	}
	return e
}
