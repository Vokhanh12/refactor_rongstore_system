package apperrors

import "fmt"

type AppError struct {
	Code         string
	Status       int
	GRPCCode     string
	Key          string
	Cause        string
	ClientAction string
	ServerAction string
	Source       string
	Component    string
	Tags         []string
	Message      string
	Data         map[string]interface{}
	Severity     string
	Expected     bool
	Retryable    bool
	CauseDetail  error
	ErrorDetails *[]AppErrorDetail
}

func (e *AppError) Error() string {
	if e == nil {
		return "<nil AppError>"
	}
	return fmt.Sprintf("%s (%d): %s", e.Code, e.Status, e.Message)
}
