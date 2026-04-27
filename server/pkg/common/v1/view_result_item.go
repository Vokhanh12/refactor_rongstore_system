package commonv1

type ViewResultItemDTO struct {
	OpID       string            `json:"op_id,omitempty"`
	ResourceID string            `json:"resource_id,omitempty"`
	Success    bool              `json:"success,omitempty"`
	Items      any               `json:"items,omitempty"`      // google.protobuf.Any
	Error      *ErrorDTO         `json:"error,omitempty"`      // từ error.proto
	Pagination *PaginationDTO    `json:"pagination,omitempty"` // từ pagination.proto
	Warnings   []*WarningDTO     `json:"warnings,omitempty"`   // repeated Warning
	Details    map[string]string `json:"details,omitempty"`    // map<string,string>
}
