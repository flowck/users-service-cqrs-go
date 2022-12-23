package user

import (
	"errors"
)

type User struct {
	id        *ID
	firstName string
	lastName  string
	email     *Email
	isBlocked bool
}

func New(id, firstName, lastName, email string) (*User, error) {
	newId, err := NewIDFromString(id)
	if err != nil {
		return nil, err
	}

	if firstName == "" {
		return nil, errors.New("firstName cannot be empty")
	}

	if lastName == "" {
		return nil, errors.New("lastName cannot be empty")
	}

	var newEmail *Email
	if newEmail, err = NewEmail(email); err != nil {
		return nil, err
	}

	return &User{
		id:        newId,
		firstName: firstName,
		lastName:  lastName,
		email:     newEmail,
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

func (u *User) Email() *Email {
	return u.email
}

func (u *User) ID() *ID {
	return u.id
}

func (u *User) Block() {
	u.isBlocked = true
}

func (u *User) UnBlock() {
	u.isBlocked = false
}

var ErrUserNotFound = errors.New("user not found")
