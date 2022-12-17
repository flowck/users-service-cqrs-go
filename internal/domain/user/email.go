package user

import "errors"

type Email struct{ value string }

func (e *Email) String() string {
	return e.value
}

func NewEmail(value string) (Email, error) {
	if value == "" {
		return Email{}, errors.New("email cannot by empty")
	}

	return Email{value: value}, nil
}
