package apperrors

import "fmt"

type AppErrorDetail struct {
	Field   string // field bị lỗi
	Code    string // REQUIRED, INVALID_FORMAT...
	Message string // message cụ thể
	Hint    string // optional fix suggestion
}

func (e *AppErrorDetail) Error() string {
	if e == nil {
		return "<nil AppErrorDetail>"
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
