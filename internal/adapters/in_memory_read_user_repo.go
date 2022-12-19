package adapters

import (
	"context"
	"users-service-cqrs/internal/domain/user"
)

type inMemoryReadUserRepository struct {
	// data map[string]*user.User
}

func (i inMemoryReadUserRepository) FindAll(ctx context.Context) ([]*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (i inMemoryReadUserRepository) Find(ctx context.Context, id string) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewInMemoryReadUserRepository() *inMemoryReadUserRepository {
	return &inMemoryReadUserRepository{}
}
