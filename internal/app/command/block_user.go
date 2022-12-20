package command

import (
	"context"
	"users-service-cqrs/internal/common/cqrs"
	"users-service-cqrs/internal/domain/user"
)

type BlockUser struct {
	UserId *user.ID
}

type BlockUserHandler = cqrs.CommandHandler[BlockUser]

type blockUserHandler struct {
	repo user.WriteRepository
}

func NewBlockUserHandler(repo user.WriteRepository) *blockUserHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return &blockUserHandler{repo: repo}
}

func (h *blockUserHandler) Handle(ctx context.Context, cmd BlockUser) error {
	return h.repo.Update(ctx, cmd.UserId, func(u *user.User) *user.User {
		u.Block()
		return u
	})
}
