package dispatcher

type Error struct {
	OpID string
	Err  error
}

type Errors []Error

func (e Errors) Error() string {
	return "multiple errors"
}
