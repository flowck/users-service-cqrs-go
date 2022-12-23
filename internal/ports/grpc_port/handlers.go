package grpc_port

import (
	"context"
	"users-service-cqrs/internal/app"
	"users-service-cqrs/internal/domain/user"
)

type Handlers struct {
	UnimplementedUsersServiceServer
	application *app.App
}

func (h Handlers) BlockUser(ctx context.Context, request *BlockUserRequest) (*Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handlers) UnblockUser(ctx context.Context, request *UnBlockUserRequest) (*Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handlers) GetAllUsers(ctx context.Context, request *GetAllUsersRequest) (*GetAllUsersResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handlers) GetOneUser(ctx context.Context, request *GetOneUserRequest) (*User, error) {
	id, err := user.NewIDFromString(request.Id)
	if err != nil {
		return nil, err
	}

	u, err := h.application.Queries.OneUser.Handle(ctx, &id)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:        u.Id,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Status:    getUserStatus(u.Status),
	}, nil
}

func getUserStatus(status string) UserStatus {
	if status == "blocked" {
		return UserStatus_blocked
	}

	return UserStatus_unblocked
}

var _ UsersServiceServer = (*Handlers)(nil)
