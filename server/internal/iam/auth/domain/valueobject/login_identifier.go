package valueobject

import (
	"net/mail"
	"strings"

	"github.com/vokhanh12/refactor-rongstore-system/server/internal/core/domain/validator"
)

type LoginIdentifierType string

const (
	LoginIdentifierEmail LoginIdentifierType = "EMAIL"
	LoginIdentifierPhone LoginIdentifierType = "PHONE"
)

type LoginIdentifier struct {
	value string
	kind  LoginIdentifierType
}

func NewLoginIdentifier(value string) (LoginIdentifier, error) {
	value = strings.TrimSpace(value)

	err := validator.New().
		Required("identifier", value).
		EmailOrPhone("identifier", value).
		Err()

	if err != nil {
		return LoginIdentifier{}, err
	}

	kind := detectLoginIdentifierType(value)

	// Chuẩn hóa giá trị theo loại identifier
	if kind == LoginIdentifierEmail {
		value = strings.ToLower(value) // Đưa email về chữ thường
	} else {
		// value = normalizePhoneNumber(value) // Chuẩn hóa số điện thoại về dạng E.164 nếu cần
	}

	return LoginIdentifier{
		value: value,
		kind:  kind,
	}, nil
}

func detectLoginIdentifierType(value string) LoginIdentifierType {
	if _, err := mail.ParseAddress(value); err == nil {
		return LoginIdentifierEmail
	}

	return LoginIdentifierPhone
}

func (i LoginIdentifier) Value() string {
	return i.value
}

func (i LoginIdentifier) Kind() LoginIdentifierType {
	return i.kind
}
