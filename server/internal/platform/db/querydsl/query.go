package querydsl

type Query struct {
	Keyword string

	Filters []Filter
	Sorts   []Sort

	Pagination *Pagination
}
