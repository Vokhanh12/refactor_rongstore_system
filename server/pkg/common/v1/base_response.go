package commonv1

import (
	"github.com/golang/protobuf/ptypes/any"
)

// BaseResponse tương ứng với proto BaseResponse
type BaseResponse struct {
	Success    bool              `json:"success"`
	Data       *any.Any          `json:"data,omitempty"`       // protobuf Any
	Metadata   *Metadata         `json:"metadata,omitempty"`   // import từ metadata.proto
	Error      *Error            `json:"error,omitempty"`      // import từ error.proto
	Pagination *Pagination       `json:"pagination,omitempty"` // import từ pagination.proto
	Warnings   []*Warning        `json:"warnings,omitempty"`   // import từ warning.proto
	Details    map[string]string `json:"details,omitempty"`    // key-value details
}

// MutateResponse tương ứng với proto MutateResponse
type MutateResponse struct {
	Metadata      *Metadata       `json:"metadata,omitempty"`
	MutateResults []*MutateResult `json:"mutate_results,omitempty"` // import từ mutate_result.proto
}

// ViewResponse tương ứng với proto ViewResponse
type ViewResponse struct {
	Metadata    *Metadata     `json:"metadata,omitempty"`
	ViewResults []*ViewResult `json:"view_results,omitempty"` // import từ view_result.proto
}

// Placeholder structs cho các proto import
// Bạn nên replace bằng struct generate từ proto thực tế
type Metadata struct {
	RequestID string            `json:"request_id"`
	Tags      map[string]string `json:"tags,omitempty"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Pagination struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`
	Total    int32 `json:"total"`
}

type Warning struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type MutateResult struct {
	ID      string `json:"id"`
	Success bool   `json:"success"`
}

type ViewResult struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}
