package commonv1

// BaseRequest tương ứng với proto BaseRequest
type BaseRequest struct {
	RequestDateTime string `json:"request_date_time" validate:"required,datetime=2006-01-02T15:04:05.00Z"` // pattern ISO8601
	RequestID       string `json:"request_id" validate:"required,uuid"`                                    // uuid validate
	Language        string `json:"language" validate:"required"`                                           // min_len:1
	UserAgent       string `json:"user_agent" validate:"required"`                                         // min_len:1
}

// SearchRequest tương ứng với proto SearchRequest
type SearchRequest struct {
	Pagination *Pagination `json:"pagination" validate:"required"` // proto message validate required
	Input      *Search     `json:"input" validate:"required"`      // proto message validate required
}

// Pagination struct placeholder (proto import common/v1/pagination.proto)
type Pagination struct {
	Page     int32 `json:"page"`
	PageSize int32 `json:"page_size"`
}

// Search struct placeholder (proto import common/v1/search.proto)
type Search struct {
	Query string `json:"query"`
}
