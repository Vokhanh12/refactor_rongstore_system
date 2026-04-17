package commonv1

type MutateResultItem struct {
	OpID    string    `json:"op_id,omitempty"`
	Data    any       `json:"data,omitempty"`
	Success bool      `json:"success,omitempty"`
	Error   *AppError `json:"error,omitempty"`
}

func NewMutateResultItem(opID string, data any, aerr AppError) *MutateResultItem {
	return &MutateResultItem{
		OpID:    opID,
		Data:    data,
		Success: false,
		Error:   &aerr,
	}
}
