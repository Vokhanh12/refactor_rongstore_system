package apperrors

import "fmt"

type Violation struct {
	Field   string // field bị lỗi
	Code    string // REQUIRED, INVALID_FORMAT...
	Message string // message cụ thể
	Hint    string // optional fix suggestion
}

func (e *Violation) Error() string {
	if e == nil {
		return "<nil Violation>"
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
