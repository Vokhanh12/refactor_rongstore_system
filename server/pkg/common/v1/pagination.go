package commonv1

// Pagination tương ứng với proto Pagination
type Pagination struct {
	Limit  int32 `json:"limit,omitempty" validate:"gt=0"`   // phải > 0
	Offset int32 `json:"offset,omitempty" validate:"gte=0"` // >= 0
}
