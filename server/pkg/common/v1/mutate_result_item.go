package commonv1

type MutateResultItemDTO struct {
	OpID    string               `json:"op_id,omitempty"`
	Data    any                  `json:"data,omitempty"`
	Success bool                 `json:"success,omitempty"`
	Error   *ExternalAppErrorDTO `json:"error,omitempty"`
}

func NewMutateResultItemDTO(opID string, data any, aerr ExternalAppErrorDTO) *MutateResultItemDTO {
	return &MutateResultItemDTO{
		OpID:    opID,
		Data:    data,
		Success: false,
		Error:   &aerr,
	}
}
