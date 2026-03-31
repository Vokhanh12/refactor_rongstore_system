package commonv1

import (
	"github.com/golang/protobuf/ptypes/any"
)

type BaseResponse struct {
	Success    bool              `json:"success"`
	Data       *any.Any          `json:"data,omitempty"`       // protobuf Any
	Metadata   *Metadata         `json:"metadata,omitempty"`   // import từ metadata.proto
	Error      *AppError         `json:"error,omitempty"`      // import từ error.proto
	Pagination *Pagination       `json:"pagination,omitempty"` // import từ pagination.proto
	Warnings   []*Warning        `json:"warnings,omitempty"`   // import từ warning.proto
	Details    map[string]string `json:"details,omitempty"`    // key-value details
}

type MutateResponse struct {
	Metadata      *Metadata       `json:"metadata,omitempty"`
	MutateResults []*MutateResult `json:"mutate_results,omitempty"` // import từ mutate_result.proto
}

type ViewResponse struct {
	Metadata    *Metadata     `json:"metadata,omitempty"`
	ViewResults []*ViewResult `json:"view_results,omitempty"` // import từ view_result.proto
}
