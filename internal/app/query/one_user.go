package query

import (
	"context"
	"users-service-cqrs/internal/common/cqrs"
	"users-service-cqrs/internal/domain/user"
)

type oneUserHandler struct {
	repo ReadRepository
}

type OneUserHandler = cqrs.Query[*user.ID, *User]

func NewOneUserHandler(repo ReadRepository) *oneUserHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return &oneUserHandler{repo: repo}
}

func (o oneUserHandler) Handle(ctx context.Context, id *user.ID) (*User, error) {
	return o.repo.Find(ctx, id)
}
