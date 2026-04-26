package repositories

type ExistsRecord struct {
	Empty bool
	Err   error
}

type CreateRecord[T any] struct {
	Obj T
	Err error
}

type UpdateRecord[T any] struct {
	Obj T
	Err error
}

type DeleteRecord[T any] struct {
	Obj T
	Err error
}
