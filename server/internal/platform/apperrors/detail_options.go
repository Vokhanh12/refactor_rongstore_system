package apperrors

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

func WithHint(hint string) func(*AppErrorDetail) {
	return func(e *AppErrorDetail) {
		e.Hint = hint
	}
}
