package querydsl

type FieldType string

const (
	FieldTypeString FieldType = "STRING"
	FieldTypeBool   FieldType = "BOOL"
	FieldTypeInt    FieldType = "INT"
	FieldTypeTime   FieldType = "TIME"
)

type Field struct {
	Column string
	Type   FieldType

	Searchable bool
	Sortable   bool
	Filterable bool

	AllowedOperators []FilterOperator
}
