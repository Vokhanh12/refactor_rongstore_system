package dispatcher

type MultipleError struct {
	Errors []*OperationError
}

func (e *MultipleError) Add(opID string, err error) {
	if err == nil {
		return
	}

	e.Errors = append(e.Errors, &OperationError{
		OpID: opID,
		Err:  err,
	})
}

func (e *MultipleError) Empty() bool {
	return len(e.Errors) == 0
}

func (e *MultipleError) Error() string {
	return "multiple operation errors"
}

// ============================================================
// Error information
// ============================================================

func (e *MultipleError) Total() int {
	if e == nil {
		return 0
	}

	return len(e.Errors)
}
