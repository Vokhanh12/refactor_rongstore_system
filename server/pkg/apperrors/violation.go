package apperrors

import "fmt"

type Violation struct {
	Field   string `json:"field"`          // field bị lỗi
	Code    string `json:"code"`           // REQUIRED, INVALID_FORMAT...
	Message string `json:"message"`        // message cụ thể
	Hint    string `json:"hint,omitempty"` // optional fix suggestion
}

func (e *Violation) Error() string {
	if e == nil {
		return "<nil Violation>"
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
