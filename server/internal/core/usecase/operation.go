package usecases

type Operation[T any] struct {
	OpID    string
	Payload T
}
