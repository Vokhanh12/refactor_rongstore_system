package commonv1

type ExternalAppErrorDetailDTO struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
	Hint    string `json:"hint,omitempty"`
}

type ExternalAppErrorDTO struct {
	Code    string                       `json:"code,omitempty"`
	Message string                       `json:"message,omitempty"`
	Details *[]ExternalAppErrorDetailDTO `json:"details,omitempty"`
}

func NewExternalAppErrorDTO(code string, msg string, aerrdetail *[]ExternalAppErrorDetailDTO) *ExternalAppErrorDTO {
	return &ExternalAppErrorDTO{
		Message: msg,
		Code:    code,
		Details: aerrdetail,
	}
}

func NewExternalAppErrorDetailDTO(code string, msg string, field string, hint string) *ExternalAppErrorDetailDTO {
	return &ExternalAppErrorDetailDTO{
		Field:   field,
		Message: msg,
		Code:    code,
		Hint:    hint,
	}
}
