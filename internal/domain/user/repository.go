package user

import "context"

type ReadRepository interface {
	FindAll(ctx context.Context) ([]*User, error)
	Find(ctx context.Context, id *ID) (*User, error)
}

type WriteRepository interface {
	Update(ctx context.Context, u *User) error
}
