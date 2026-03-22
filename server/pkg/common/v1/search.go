package commonv1

// Search tương ứng với proto Search
type Search struct {
	Keyword string `json:"keyword,omitempty" validate:"required,min=1"` // min_len:1
}
