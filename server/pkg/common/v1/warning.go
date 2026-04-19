package commonv1

type WarningDTO struct {
	Code     string            `json:"code,omitempty"`
	Message  string            `json:"message,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"` // key-value metadata
}
