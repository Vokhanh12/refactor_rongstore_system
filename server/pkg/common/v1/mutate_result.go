package commonv1

import (
	"github.com/golang/protobuf/ptypes/any"
)

// MutateResult tương ứng với proto MutateResult
type MutateResult struct {
	OpID            string              `json:"op_id,omitempty"`
	Data            *any.Any            `json:"data,omitempty"` // google.protobuf.Any
	Success         bool                `json:"success,omitempty"`
	MutateOperation MutateOperationEnum `json:"mutate_operation,omitempty"` // từ MutateOperation.proto
	Error           *Error              `json:"error,omitempty"`            // từ error.proto
}
