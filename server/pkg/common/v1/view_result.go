package commonv1

import (
	"github.com/golang/protobuf/ptypes/any"
)

type ViewResult struct {
	OpID       string            `json:"op_id,omitempty"`
	ResourceID string            `json:"resource_id,omitempty"`
	Success    bool              `json:"success,omitempty"`
	Items      *any.Any          `json:"items,omitempty"`      // google.protobuf.Any
	Error      *AppError         `json:"error,omitempty"`      // từ error.proto
	Pagination *Pagination       `json:"pagination,omitempty"` // từ pagination.proto
	Warnings   []*Warning        `json:"warnings,omitempty"`   // repeated Warning
	Details    map[string]string `json:"details,omitempty"`    // map<string,string>
}
