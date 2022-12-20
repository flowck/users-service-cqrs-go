package query

import (
	"context"
	"users-service-cqrs/internal/common/cqrs"
)

type AllUsers struct {
	Status string
}

type AllUsersHandler = cqrs.Query[AllUsers, []*User]

type allUsersHandler struct {
	repo ReadRepository
}

func NewAllBlockedUsersHandler(repo ReadRepository) *allUsersHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return &allUsersHandler{repo: repo}
}

func (h allUsersHandler) Handle(ctx context.Context, q AllUsers) ([]*User, error) {
	return h.repo.FindAll(ctx, q)
}
