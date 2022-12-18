package user

import (
	"errors"
	"regexp"
)

type Email struct{ value string }

func (e *Email) String() string {
	return e.value
}

func NewEmail(value string) (Email, error) {
	regex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if value == "" {
		return Email{}, ErrEmailEmpty
	}

	if !regex.MatchString(value) {
		return Email{}, ErrEmailInvalid
	}

	return Email{value: value}, nil
}

// Errors

var ErrEmailInvalid = errors.New("email seems to be invalid")
var ErrEmailEmpty = errors.New("email seems is empty")
