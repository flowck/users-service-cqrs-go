package command

import (
	"context"
	"users-service-cqrs/internal/common/cqrs"
	"users-service-cqrs/internal/domain/user"
)

type UnblockUser struct {
	UserId string
}

type UnBlockUserHandler = cqrs.CommandHandler[UnblockUser]

type unblockUserHandler struct {
	repo user.WriteRepository
}

func NewUnblockUserHandler(repo user.WriteRepository) *unblockUserHandler {
	if repo == nil {
		panic("repo is nil")
	}

	return &unblockUserHandler{repo: repo}
}

func (h *unblockUserHandler) Handle(ctx context.Context, cmd UnblockUser) error {
	return nil
}
