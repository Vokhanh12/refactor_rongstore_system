package apperrors

import "fmt"

type AppError struct {
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
	causeDetail  error
	ErrorDetails *[]AppErrorDetail
}

func (e *AppError) Error() string {
	if e == nil {
		return "<nil AppError>"
	}
	return fmt.Sprintf("%s (%d): %s", e.Code, e.Status, e.Message)
}

func (e *AppError) GetCauseDetail() error {
	if e == nil {
		return nil
	}
	return e.causeDetail
}

func (e *AppError) GetErrorDetails() []AppErrorDetail {
	if e == nil {
		return nil
	}
	return e.errorDetails
}
