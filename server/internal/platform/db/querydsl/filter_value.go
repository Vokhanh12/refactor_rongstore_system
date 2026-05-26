package querydsl

func StringValue(
	v *FilterValue,
) (string, bool) {

	x, ok := v.Kind.(FilterValueString)

	if !ok {
		return "", false
	}

	return x.Value, true
}

func IntValue(
	v *FilterValue,
) (int64, bool) {

	x, ok := v.Kind.(FilterValueInt)

	if !ok {
		return 0, false
	}

	return x.Value, true
}

func BoolValue(
	v *FilterValue,
) (bool, bool) {

	x, ok := v.Kind.(FilterValueBool)

	if !ok {
		return false, false
	}

	return x.Value, true
}

func ParseFilterValue(
	field Field,
	value *FilterValue,
) (any, bool) {

	switch field.Type {

	case FieldTypeString:
		return StringValue(value)

	case FieldTypeInt:
		return IntValue(value)

	case FieldTypeBool:
		return BoolValue(value)
	}

	return nil, false
}

func ParseFilterValues(
	field Field,
	values []*FilterValue,
) ([]any, bool) {

	result := make([]any, 0, len(values))

	for _, value := range values {

		v, ok := ParseFilterValue(
			field,
			value,
		)

		if !ok {
			return nil, false
		}

		result = append(result, v)
	}

	return result, true
}
