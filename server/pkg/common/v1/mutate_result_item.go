package commonv1

type MutateResultItemDTO struct {
	OpID    string    `json:"op_id,omitempty"`
	Data    any       `json:"data,omitempty"`
	Success bool      `json:"success,omitempty"`
	Error   *ErrorDTO `json:"error,omitempty"`
}
