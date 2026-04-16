package commonv1

type AppErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
	Hint    string `json:"hint,omitempty"`
}

type AppError struct {
	Code         string           `json:"code,omitempty"`        // AUTH-HAND-001
	Key          string           `json:"key,omitempty"`         // HANDSHAKE_INVALID_CLIENT_KEY
	Message      string           `json:"message,omitempty"`     // "Invalid client public key"
	Severity     string           `json:"severity,omitempty"`    // S1/S2/S3
	Retryable    bool             `json:"retryable,omitempty"`   // true/false
	Source       string           `json:"source,omitempty"`      // iam-service, gateway,...
	GRPCCode     string           `json:"grpc_code,omitempty"`   // InvalidArgument
	HTTPStatus   int32            `json:"http_status,omitempty"` // 400, 401, 500...
	ClientAction string           `json:"client_action,omitempty"`
	ServerAction string           `json:"server_action,omitempty"`
	Details      []AppErrorDetail `json:"details,omitempty"` // repeated ErrorDetail
}
