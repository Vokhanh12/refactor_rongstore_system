package valueobjects

import (
	"errors"
	"regexp"
)

type PhoneNumber struct {
	value string
}

func NewPhoneNumber(phone string) (PhoneNumber, error) {
	re := regexp.MustCompile(`^\+?[0-9]{7,15}$`)
	if !re.MatchString(phone) {
		return PhoneNumber{}, errors.New("invalid phone number")
	}
	return PhoneNumber{value: phone}, nil
}

func (p PhoneNumber) String() string {
	return p.value
}
