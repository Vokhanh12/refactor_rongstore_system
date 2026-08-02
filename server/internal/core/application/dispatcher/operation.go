package dispatcher

type Action string

type Operation struct {
	OpID    string
	Action  Action
	Payload any
}
