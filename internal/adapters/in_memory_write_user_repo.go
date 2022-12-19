package adapters

import (
	"context"
	"users-service-cqrs/internal/domain/user"
)

type inMemoryWriteUserRepository struct {
	// data map[string]*user.User
}

func (i inMemoryWriteUserRepository) Update(ctx context.Context, u *user.User) error {
	//TODO implement me
	panic("implement me")
}

func NewInMemoryWriteUserRepository() *inMemoryWriteUserRepository {
	return &inMemoryWriteUserRepository{}
}
