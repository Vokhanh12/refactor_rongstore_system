package querydsl

func IsOperatorAllowed(
	field Field,
	operator FilterOperator,
) bool {

	for _, op := range field.AllowedOperators {

		if op == operator {
			return true
		}
	}

	return false
}
