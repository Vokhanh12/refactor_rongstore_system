package fields

import "github.com/vokhanh12/refactor-rongstore-system/server/internal/platform/db/querydsl"

var RoleFields = map[string]querydsl.Field{

	"id": {
		Column: "r.id",
		Type:   querydsl.FieldTypeString,
	},

	"name": {
		Column: "r.name",
		Type:   querydsl.FieldTypeString,

		Searchable: true,
		Sortable:   true,
		Filterable: true,

		AllowedOperators: []querydsl.FilterOperator{
			querydsl.FilterOperator_EQ,
			querydsl.FilterOperator_LIKE,
			querydsl.FilterOperator_IN,
		},
	},

	"is_active": {
		Column: "r.is_active",
		Type:   querydsl.FieldTypeBool,

		Filterable: true,

		AllowedOperators: []querydsl.FilterOperator{
			querydsl.FilterOperator_EQ,
			querydsl.FilterOperator_NEQ,
		},
	},

	"level": {
		Column: "r.level",
		Type:   querydsl.FieldTypeInt,

		Filterable: true,
		Sortable:   true,

		AllowedOperators: []querydsl.FilterOperator{
			querydsl.FilterOperator_EQ,
			querydsl.FilterOperator_GT,
			querydsl.FilterOperator_GTE,
			querydsl.FilterOperator_LT,
			querydsl.FilterOperator_LTE,
		},
	},
}
