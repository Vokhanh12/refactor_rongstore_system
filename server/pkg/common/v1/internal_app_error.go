package commonv1

type InternalAppErrorDTO struct {
	Code         string                 `json:"code"`
	Status       int                    `json:"status"`
	GRPCCode     string                 `json:"grpc_code"`
	Key          string                 `json:"key"`
	Cause        string                 `json:"cause"`
	ClientAction string                 `json:"client_action"`
	ServerAction string                 `json:"server_action"`
	Source       string                 `json:"source"`
	Component    string                 `json:"component"`
	Tags         []string               `json:"tags"`
	Message      string                 `json:"message"`
	Data         map[string]interface{} `json:"data,omitempty"`
	Severity     string                 `json:"severity,omitempty"`
	Expected     bool                   `json:"expected,omitempty"`
	Retryable    bool                   `json:"retryable,omitempty"`
	CauseDetail  string                 `json:"cause_detail,omitempty"`
	ErrorDetails *[]InternalAppErrorDetailDTO
}

type InternalAppErrorDetailDTO struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
	Hint    string `json:"hint,omitempty"`
}

func NewInternalAppErrorDTO(code string, msg string, aerrdetail *[]InternalAppErrorDetailDTO) *InternalAppErrorDTO {
	return &InternalAppErrorDTO{
		Message:      msg,
		Code:         code,
		ErrorDetails: aerrdetail,
	}
}

func NewInternalAppErrorDetailDTO(code string, msg string, field string, hint string) *InternalAppErrorDetailDTO {
	return &InternalAppErrorDetailDTO{
		Field:   field,
		Message: msg,
		Code:    code,
		Hint:    hint,
	}
}
