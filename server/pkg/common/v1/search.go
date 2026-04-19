package commonv1

type SearchDTO struct {
	Keyword string `json:"keyword,omitempty" validate:"required,min=1"` // min_len:1
}
