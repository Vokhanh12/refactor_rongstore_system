package commonv1

type MutateResultItem struct {
	OpID    string    `json:"op_id,omitempty"`
	Data    any       `json:"data,omitempty"`
	Success bool      `json:"success,omitempty"`
	Code    string    `json:"code,omitempty"`
	Error   *AppError `json:"error,omitempty"`
}
