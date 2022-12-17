package query

import (
	"context"
	"users-service-cqrs/internal/common/cqrs"
	"users-service-cqrs/internal/domain/user"
)

type AllBlockedUsers struct{}

type AllBlockedUsersHandler = cqrs.Query[AllBlockedUsers, []*User]

type allBlockedUsersHandler struct {
	repo user.ReadRepository
}

func NewAllBlockedUsersHandler(repo user.ReadRepository) *allBlockedUsersHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return &allBlockedUsersHandler{repo: repo}
}

func (h allBlockedUsersHandler) Handle(ctx context.Context, q AllBlockedUsers) ([]*User, error) {
	return nil, nil
}
