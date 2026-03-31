package commonv1

type Warning struct {
	Code     string            `json:"code,omitempty"`
	Message  string            `json:"message,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"` // key-value metadata
}
