package commonv1

type ViewResultItemDTO struct {
	OpID    string       `json:"op_id,omitempty"`
	Data    any          `json:"data,omitempty"`
	Success bool         `json:"success,omitempty"`
	Code    string       `json:"code,omitempty"`
	Error   *AppErrorDTO `json:"error,omitempty"`
}
