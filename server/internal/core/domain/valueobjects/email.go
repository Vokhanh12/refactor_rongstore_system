package valueobjects

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	if !re.MatchString(email) {
		return Email{}, errors.New("invalid email")
	}
	return Email{value: email}, nil
}

func (e Email) String() string {
	return e.value
}
