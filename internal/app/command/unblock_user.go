package command

import (
	"context"
	"users-service-cqrs/internal/common/cqrs"
	"users-service-cqrs/internal/domain/user"
)

type UnBlockUser struct {
	UserId *user.ID
}

type UnBlockUserHandler = cqrs.CommandHandler[UnBlockUser]

type unblockUserHandler struct {
	repo user.WriteRepository
}

func NewUnblockUserHandler(repo user.WriteRepository) *unblockUserHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return &unblockUserHandler{repo: repo}
}

func (h *unblockUserHandler) Handle(ctx context.Context, cmd UnBlockUser) error {
	return h.repo.Update(ctx, cmd.UserId, func(u *user.User) *user.User {
		u.UnBlock()
		return u
	})
}
