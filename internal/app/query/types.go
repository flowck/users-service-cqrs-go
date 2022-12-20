package query

import (
	"context"
	"errors"
	"users-service-cqrs/internal/domain/user"
)

// Types at the application layer
// For instance, the User domain struct might contain the password field whereas the User here doesn't need to
// if the password is not needed to realise the use cases.

type ReadRepository interface {
	FindAll(ctx context.Context, q AllUsers) ([]*User, error)
	Find(ctx context.Context, id *user.ID) (*User, error)
}

var ErrUserNotFound = errors.New("user not found")

type User struct {
	Id        string
	FirstName string
	LastName  string
	Email     string
	Status    string
}
