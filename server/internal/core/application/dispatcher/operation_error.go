package dispatcher

type OperationError struct {
	OpID string
	Err  error
}

func (e *OperationError) Error() string {
	return e.Err.Error()
}
