package commonv1

type BaseRequest struct {
	RequestDateTime string `json:"request_date_time" validate:"required,datetime=2006-01-02T15:04:05.00Z"` // pattern ISO8601
	RequestID       string `json:"request_id" validate:"required,uuid"`                                    // uuid validate
	Language        string `json:"language" validate:"required"`                                           // min_len:1
	UserAgent       string `json:"user_agent" validate:"required"`                                         // min_len:1
}

type SearchRequest struct {
	Pagination *Pagination `json:"pagination" validate:"required"` // proto message validate required
	Input      *Search     `json:"input" validate:"required"`      // proto message validate required
}
