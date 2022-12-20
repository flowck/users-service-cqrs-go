package adapters

import (
	"context"
	"users-service-cqrs/internal/domain/user"
)

type inMemoryWriteUserRepository struct {
	// data map[string]*user.User
}

var _ user.WriteRepository = (*inMemoryWriteUserRepository)(nil)

func (i inMemoryWriteUserRepository) Update(ctx context.Context, id *user.ID, updateFn func(u *user.User) *user.User) error {
	//TODO implement me
	panic("implement me")
}

func NewInMemoryWriteUserRepository() *inMemoryWriteUserRepository {
	return &inMemoryWriteUserRepository{}
}
