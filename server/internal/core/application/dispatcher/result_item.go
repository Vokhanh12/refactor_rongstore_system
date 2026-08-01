package dispatcher

import aerrs "github.com/vokhanh12/refactor-rongstore-system/server/pkg/apperrors"

type ResultItem struct {
	OpID       string            `json:"op_id,omitempty"`
	ResourceID string            `json:"resource_id,omitempty"`
	Data       any               `json:"data,omitempty"`
	Success    bool              `json:"success,omitempty"`
	Error      *aerrs.AppError   `json:"error,omitempty"`
	Details    map[string]string `json:"details,omitempty"`
}
