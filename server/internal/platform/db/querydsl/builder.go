package querydsl

import (
	"strings"

	sq "github.com/Masterminds/squirrel"
)

type Builder struct {
	fields map[string]Field
}

func NewBuilder(
	fields map[string]Field,
) *Builder {

	return &Builder{
		fields: fields,
	}
}

func (b *Builder) ApplySearch(
	qb sq.SelectBuilder,
	keyword string,
) sq.SelectBuilder {

	if keyword == "" {
		return qb
	}

	conditions := sq.Or{}

	for _, field := range b.fields {

		if !field.Searchable {
			continue
		}

		conditions = append(
			conditions,
			sq.ILike{
				field.Column: "%" + keyword + "%",
			},
		)
	}

	if len(conditions) == 0 {
		return qb
	}

	return qb.Where(conditions)
}

func (b *Builder) ApplyFilters(
	qb sq.SelectBuilder,
	filters []Filter,
) sq.SelectBuilder {

	for _, filter := range filters {

		field, ok := b.fields[filter.Field]

		if !ok {
			continue
		}

		if !field.Filterable {
			continue
		}

		if !IsOperatorAllowed(field, filter.Op) {
			continue
		}

		if len(filter.Values) == 0 {
			continue
		}

		switch filter.Op {

		//
		// EQ
		//

		case FilterOperator_EQ:

			value, ok := ParseFilterValue(
				field,
				filter.Values[0],
			)

			if !ok {
				continue
			}

			qb = qb.Where(
				sq.Eq{
					field.Column: value,
				},
			)

		//
		// NEQ
		//

		case FilterOperator_NEQ:

			value, ok := ParseFilterValue(
				field,
				filter.Values[0],
			)

			if !ok {
				continue
			}

			qb = qb.Where(
				sq.NotEq{
					field.Column: value,
				},
			)

		//
		// GT
		//

		case FilterOperator_GT:

			value, ok := ParseFilterValue(
				field,
				filter.Values[0],
			)

			if !ok {
				continue
			}

			qb = qb.Where(
				sq.Gt{
					field.Column: value,
				},
			)

		//
		// GTE
		//

		case FilterOperator_GTE:

			value, ok := ParseFilterValue(
				field,
				filter.Values[0],
			)

			if !ok {
				continue
			}

			qb = qb.Where(
				sq.GtOrEq{
					field.Column: value,
				},
			)

		//
		// LT
		//

		case FilterOperator_LT:

			value, ok := ParseFilterValue(
				field,
				filter.Values[0],
			)

			if !ok {
				continue
			}

			qb = qb.Where(
				sq.Lt{
					field.Column: value,
				},
			)

		//
		// LTE
		//

		case FilterOperator_LTE:

			value, ok := ParseFilterValue(
				field,
				filter.Values[0],
			)

			if !ok {
				continue
			}

			qb = qb.Where(
				sq.LtOrEq{
					field.Column: value,
				},
			)

		//
		// LIKE
		//

		case FilterOperator_LIKE:

			value, ok := ParseFilterValue(
				field,
				filter.Values[0],
			)

			if !ok {
				continue
			}

			s, ok := value.(string)

			if !ok {
				continue
			}

			qb = qb.Where(
				sq.ILike{
					field.Column: "%" + s + "%",
				},
			)

		//
		// IN
		//

		case FilterOperator_IN:

			values, ok := ParseFilterValues(
				field,
				filter.Values,
			)

			if !ok {
				continue
			}

			qb = qb.Where(
				sq.Eq{
					field.Column: values,
				},
			)
		}
	}

	return qb
}

func (b *Builder) ApplySorts(
	qb sq.SelectBuilder,
	sorts []Sort,
) sq.SelectBuilder {

	for _, sort := range sorts {

		field, ok := b.fields[sort.Field]

		if !ok {
			continue
		}

		if !field.Sortable {
			continue
		}

		direction := "ASC"

		if sort.Direction == SortDirection_DESC {
			direction = "DESC"
		}

		qb = qb.OrderBy(
			strings.Join(
				[]string{
					field.Column,
					direction,
				},
				" ",
			),
		)
	}

	return qb
}

func (b *Builder) ApplyPagination(
	qb sq.SelectBuilder,
	pagination *Pagination,
) sq.SelectBuilder {

	if pagination == nil {
		return qb
	}

	return qb.
		Limit(pagination.Limit).
		Offset(pagination.Offset)
}
