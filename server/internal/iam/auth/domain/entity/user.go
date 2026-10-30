package entity

type UserStatus string

const (
	StatusActive   UserStatus = "ACTIVE"
	StatusInactive UserStatus = "INACTIVE"
	StatusBlocked  UserStatus = "BLOCKED"
)

type User struct {
	id     string
	status UserStatus
}

func NewUser(id string) User {
	return User{
		id:     id,
		status: StatusActive,
	}
}

func (u User) ID() string {
	return u.id
}

func (u User) Status() UserStatus {
	return u.status
}

func (u *User) Block() {
	u.status = StatusBlocked
}
