package commonv1

type BaseRequestDTO struct {
	RequestDateTime string `json:"request_date_time" validate:"required,datetime=2006-01-02T15:04:05.00Z"`
	RequestID       string `json:"request_id" validate:"required,uuid"`
	Language        string `json:"language" validate:"required"`
	UserAgent       string `json:"user_agent" validate:"required"`
}

type SearchRequestDTO struct {
	Pagination *Pagination `json:"pagination" validate:"required"`
	Input      *Search     `json:"input" validate:"required"`
}
