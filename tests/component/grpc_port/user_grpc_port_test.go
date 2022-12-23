package grpc_port_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	pb "users-service-cqrs/internal/ports/grpc_port"
)

func TestGetOneUser(t *testing.T) {
	service, conn := createService()
	ctx, cancel := context.WithCancel(context.Background())
	defer conn.Close()
	defer cancel()

	t.Run("Expect user to be returned", func(t *testing.T) {
		id := "246cf82b-f50e-4ee6-b948-ccfe938b7d2f"
		u, err := service.GetOneUser(ctx, &pb.GetOneUserRequest{Id: id})
		assert.Nil(t, err)
		assertUser(t, u)
		assert.Equal(t, u.Id, id)
	})

	t.Run("Expect user to not be returned when the id provided isn't in the db", func(t *testing.T) {
		u, err := service.GetOneUser(ctx, &pb.GetOneUserRequest{Id: "246cf82b-f50e-4ee6-b948-ccfe938b7d2t"})
		assert.Nil(t, u)
		assert.NotNil(t, err)
	})
}

func createService() (pb.UsersServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost:3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return pb.NewUsersServiceClient(conn), conn
}

func assertUser(t *testing.T, u *pb.User) {
	assert.NotEmpty(t, u.Email)
	assert.NotEmpty(t, u.FirstName)
	assert.NotEmpty(t, u.LastName)
	assert.NotEmpty(t, u.Status.String())
	assert.NotEmpty(t, u.Id)
}
