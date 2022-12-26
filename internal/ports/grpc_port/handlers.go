package grpc_port

import (
	"context"
	"users-service-cqrs/internal/app"
	"users-service-cqrs/internal/app/command"
	"users-service-cqrs/internal/app/query"
	"users-service-cqrs/internal/domain/user"
)

type Handlers struct {
	UnimplementedUsersServiceServer
	application *app.App
}

func (h Handlers) BlockUser(ctx context.Context, req *BlockUserRequest) (*Empty, error) {
	id, err := user.NewIDFromString(req.Id)
	if err != nil {
		return nil, err
	}

	if err = h.application.Commands.BlockUser.Handle(ctx, command.BlockUser{UserId: id}); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h Handlers) UnblockUser(ctx context.Context, req *UnBlockUserRequest) (*Empty, error) {
	id, err := user.NewIDFromString(req.Id)
	if err != nil {
		return nil, err
	}

	if err = h.application.Commands.UnBlockUser.Handle(ctx, command.UnBlockUser{UserId: id}); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h Handlers) GetAllUsers(ctx context.Context, req *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	status := getUserStatus(req.Status.String())
	users, err := h.application.Queries.AllUsers.Handle(ctx, query.AllUsers{Status: status.String()})
	if err != nil {
		return nil, err
	}

	userList := make([]*User, len(users))
	for idx, u := range users {
		userList[idx] = mapToProtoUser(u)
	}

	return &GetAllUsersResponse{
		Users: userList,
	}, nil
}

func (h Handlers) GetOneUser(ctx context.Context, req *GetOneUserRequest) (*User, error) {
	id, err := user.NewIDFromString(req.Id)
	if err != nil {
		return nil, err
	}

	u, err := h.application.Queries.OneUser.Handle(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapToProtoUser(u), nil
}

func mapToProtoUser(u *query.User) *User {
	return &User{
		Id:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Status:    getUserStatus(u.Status),
	}
}

func getUserStatus(status string) UserStatus {
	if status == "blocked" {
		return UserStatus_blocked
	}

	return UserStatus_unblocked
}

var _ UsersServiceServer = (*Handlers)(nil)
