package commonv1

import (
	"github.com/golang/protobuf/ptypes/any"
)

type BaseResponse struct {
	Success    bool              `json:"success"`
	Data       *any.Any          `json:"data,omitempty"`
	Metadata   *Metadata         `json:"metadata,omitempty"`
	Error      *AppError         `json:"error,omitempty"`
	Pagination *Pagination       `json:"pagination,omitempty"`
	Warnings   []*Warning        `json:"warnings,omitempty"`
	Details    map[string]string `json:"details,omitempty"`
}

type MutateResponse struct {
	Metadata      *Metadata       `json:"metadata,omitempty"`
	MutateResults []*MutateResult `json:"mutate_results,omitempty"`
}

type ViewResponse struct {
	Metadata    *Metadata     `json:"metadata,omitempty"`
	ViewResults []*ViewResult `json:"view_results,omitempty"`
}
