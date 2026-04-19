package commonv1

import (
	"github.com/golang/protobuf/ptypes/any"
)

type BaseResponseDTO struct {
	Success    bool              `json:"success"`
	Data       *any.Any          `json:"data,omitempty"`
	Metadata   *Metadata         `json:"metadata,omitempty"`
	Error      *AppErrorDTO      `json:"error,omitempty"`
	Pagination *Pagination       `json:"pagination,omitempty"`
	Warnings   []*Warning        `json:"warnings,omitempty"`
	Details    map[string]string `json:"details,omitempty"`
}

type MutateResponseDTO struct {
	Metadata      *MetadataDTO       `json:"metadata,omitempty"`
	MutateResults []*MutateResultDTO `json:"mutate_results,omitempty"`
}

type ViewResponseDTO struct {
	Metadata    *MetadataDTO     `json:"metadata,omitempty"`
	ViewResults []*ViewResultDTO `json:"view_results,omitempty"`
}
