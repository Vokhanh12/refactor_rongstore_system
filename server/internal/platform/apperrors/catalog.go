package apperrors

var catalog map[string]AppError

func InitCatalog(m map[string]AppError) {
	catalog = m
}

func Lookup(code string) *AppError {
	if e, ok := catalog[code]; ok {
		return New(e)
	}
	return New(UNKNOWN_DOMAIN_KEY)
}
