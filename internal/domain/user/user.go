package user

import "errors"

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
		Id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     Email{},
		isBlocked: false,
	}, nil
}
