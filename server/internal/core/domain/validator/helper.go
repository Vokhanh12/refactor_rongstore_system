package validator

func InEnum[T comparable](value T, set map[T]struct{}) bool {
	_, ok := set[value]
	return ok
}
