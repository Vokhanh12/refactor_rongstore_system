package commonv1

type ErrorDetailDTO struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
	Hint    string `json:"hint,omitempty"`
}

type ExternalErrorDTO struct {
	Code    string           `json:"code,omitempty"`
	Message string           `json:"message,omitempty"`
	Details []ErrorDetailDTO `json:"details"` // luôn [] thay vì nil
}

type InternalErrorDTO struct {
	Code    string `json:"code,omitempty"`    // AUTH-VAL-001
	Key     string `json:"key,omitempty"`     // LOGIN_EMAIL_EMPTY
	Message string `json:"message,omitempty"` // log friendly

	Severity  string `json:"severity,omitempty"` // S1/S2/S3
	Retryable bool   `json:"retryable,omitempty"`

	Source    string `json:"source,omitempty"`    // iam-service
	Component string `json:"component,omitempty"` // api/domain/db

	GRPCCode string `json:"grpc_code,omitempty"`

	ClientAction string `json:"client_action,omitempty"`
	ServerAction string `json:"server_action,omitempty"`
}

type ErrorDTO struct {
	External ExternalErrorDTO `json:"external,omitempty"`
	Internal InternalErrorDTO `json:"internal,omitempty"` // optional (ẩn ở production)
}
