package apperrors

func WithField(value string) func(*Violation) {
	return func(e *Violation) {
		e.Field = value
	}
}

func WithMessageDetail(msg string) func(*Violation) {
	return func(e *Violation) {
		e.Message = msg
	}
}

func WithHint(hint string) func(*Violation) {
	return func(e *Violation) {
		e.Hint = hint
	}
}
