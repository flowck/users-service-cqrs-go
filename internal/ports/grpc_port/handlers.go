package grpc_port

import (
	"context"
	"users-service-cqrs/internal/app"
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
	//TODO implement me
	panic("implement me")
}

var _ UsersServiceServer = (*Handlers)(nil)
