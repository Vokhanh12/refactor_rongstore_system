package commonv1

type AppErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
	Hint    string `json:"hint,omitempty"`
}

type AppError struct {
	Code    string           `json:"code,omitempty"`
	Message string           `json:"message,omitempty"`
	Details []AppErrorDetail `json:"details,omitempty"`
}

func NewAppError(code string, msg string, aerrdetail []AppErrorDetail) *AppError {
	return &AppError{
		Message: msg,
		Code:    code,
		Details: aerrdetail,
	}
}

func NewAppErrorDetail(code string, msg string, field string, hint string) *AppErrorDetail {
	return &AppErrorDetail{
		Field:   field,
		Message: msg,
		Code:    code,
		Hint:    hint,
	}
}
