package apperrors

import "fmt"

type AppErrorDetail struct {
	Field   string
	Message string
	Code    string // OUT_OF_RANGE, REQUIRED, NEGATIVE
}

func (e *AppErrorDetail) Error() string {
	if e == nil {
		return "<nil AppError>"
	}
	return fmt.Sprintf("%s : %s", e.Code, e.Message)
}
