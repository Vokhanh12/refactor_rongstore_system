package querydsl

type FilterOperator int32

const (
	FilterOperator_UNSPECIFIED FilterOperator = 0

	FilterOperator_EQ  FilterOperator = 1
	FilterOperator_NEQ FilterOperator = 2

	FilterOperator_GT  FilterOperator = 3
	FilterOperator_GTE FilterOperator = 4

	FilterOperator_LT  FilterOperator = 5
	FilterOperator_LTE FilterOperator = 6

	FilterOperator_LIKE FilterOperator = 7
	FilterOperator_IN   FilterOperator = 8
)

type FilterValue struct {
	Kind isFilterValueKind
}

type isFilterValueKind interface {
	isFilterValueKind()
}

//
// STRING
//

type FilterValueString struct {
	Value string
}

func (FilterValueString) isFilterValueKind() {}

//
// INT
//

type FilterValueInt struct {
	Value int64
}

func (FilterValueInt) isFilterValueKind() {}

//
// BOOL
//

type FilterValueBool struct {
	Value bool
}

func (FilterValueBool) isFilterValueKind() {}

//
// FILTER
//

type Filter struct {
	Field  string
	Op     FilterOperator
	Values []*FilterValue
}
