package commonv1

import (
	"github.com/golang/protobuf/ptypes/any"
)

// ViewResult tương ứng với proto ViewResult
type ViewResult struct {
	OpID       string            `json:"op_id,omitempty"`
	ResourceID string            `json:"resource_id,omitempty"`
	Success    bool              `json:"success,omitempty"`
	Items      *any.Any          `json:"items,omitempty"`      // google.protobuf.Any
	Error      *Error            `json:"error,omitempty"`      // từ error.proto
	Pagination *Pagination       `json:"pagination,omitempty"` // từ pagination.proto
	Warnings   []*Warning        `json:"warnings,omitempty"`   // repeated Warning
	Details    map[string]string `json:"details,omitempty"`    // map<string,string>
}
