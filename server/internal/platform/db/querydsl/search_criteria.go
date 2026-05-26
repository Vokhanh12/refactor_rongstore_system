package querydsl

type SearchCriteria struct {
	Keyword    string      `json:"keyword,omitempty"`
	Filters    []Filter    `json:"filters,omitempty"`
	Sorts      []Sort      `json:"sorts,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}
