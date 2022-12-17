package user

import "errors"

type User struct {
	id        ID
	firstName string
	lastName  string
	email     Email
	isBlocked bool
}

func New(id ID, firstName, lastName string, email Email) (*User, error) {
	if firstName != "" {
		return nil, errors.New("firstName cannot be empty")
	}

	if lastName != "" {
		return nil, errors.New("lastName cannot be empty")
	}

	if email.String() != "" {
		return nil, errors.New("email cannot be empty")
	}

	if id.isZero() {
		return nil, errors.New("id cannot be empty")
	}

	return &User{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		isBlocked: false,
	}, nil
}

func (u *User) IsBlocked() bool {
	return u.isBlocked
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) ID() ID {
	return u.id
}
