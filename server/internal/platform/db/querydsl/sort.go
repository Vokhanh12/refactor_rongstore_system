package querydsl

type SortDirection int32

const (
	SortDirection_ASC  SortDirection = 0
	SortDirection_DESC SortDirection = 1
)

type Sort struct {
	Field     string
	Direction SortDirection
}
