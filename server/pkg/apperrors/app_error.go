package apperrors

import "fmt"

type AppError struct {
	// ===== Identity =====
	Code string // AUTH-VAL-001
	Key  string // LOGIN_EMAIL_EMPTY

	// ===== Message =====
	Message string // human readable

	// ===== Protocol =====
	Status   int    // HTTP status
	GRPCCode string // gRPC status string

	// ===== Classification =====
	Component string   // api, domain, infra...
	Tags      []string // validation, auth, database...
	Severity  string   // S1, S2, S3

	// ===== Behavior =====
	Retryable bool
	Expected  bool

	// ===== Actions =====
	ClientAction string
	ServerAction string

	// ===== Debug =====
	Source      string // service name
	Cause       string // short cause (safe)
	CauseDetail error  // raw error (internal only)

	// ===== Data =====
	Data map[string]interface{}

	// ===== Validation details =====
	ErrorDetails []AppErrorDetail
}

func (e *AppError) Error() string {
	if e == nil {
		return "<nil AppError>"
	}
	return fmt.Sprintf("%s (%d): %s", e.Code, e.Status, e.Message)
}
