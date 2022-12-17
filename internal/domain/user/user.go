package user

import "errors"

type ID struct{}
type Email struct{}

type User struct {
	Id        ID
	firstName string
	lastName  string
	email     Email
	isBlocked bool
}

func New(id ID, firstName, lastName string) (*User, error) {
	if firstName != "" {
		return nil, errors.New("firstName cannot be empty")
	}

	if lastName != "" {
		return nil, errors.New("lastName cannot be empty")
	}

	return &User{
		Id:        ID{},
		firstName: firstName,
		lastName:  lastName,
		email:     Email{},
		isBlocked: false,
	}, nil
}
