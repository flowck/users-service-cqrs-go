package component

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"users-service-cqrs/tests/client"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBlockAndUnBlockUser(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cli := mustCreateClient(t)

	id, err := uuid.Parse("7e79baa9-318c-4286-9943-dd862eb65c1d")
	assert.Nil(t, err)

	t.Run("Expect user to be blocked", func(t *testing.T) {
		_, err = cli.BlockUser(ctx, id)
		assert.Nil(t, err)

		res, err := cli.GetOneUserWithResponse(ctx, id)
		assert.Equalf(t, *res.JSON200.Status, "blocked", "expect user's status %s to be blocked", id.String())
		assert.Nil(t, err)
	})

	t.Run("Expect user to be unblocked", func(t *testing.T) {
		_, err = cli.UnblockUser(ctx, id)
		assert.Nil(t, err)

		res, err := cli.GetOneUserWithResponse(ctx, id)
		assert.Equalf(t, *res.JSON200.Status, "unblocked", "expect user %s to be unblocked", id.String())
		assert.Nil(t, err)
	})
}

func TestGetOneUser(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cli := mustCreateClient(t)
	id, err := uuid.Parse("4c7e5b25-c7ac-429f-aa04-23e40a27d0e4") // From seeds
	assert.Nil(t, err)

	t.Run("Expect user to be returned", func(t *testing.T) {
		res, err := cli.GetOneUserWithResponse(ctx, id)
		assert.Nil(t, err)
		assert.Equal(t, res.JSON200.Id.String(), id.String())
		assertUser(t, res.JSON200)
	})

	t.Run("Expect to return 404 for a non-existent user", func(t *testing.T) {
		id, err = uuid.Parse("9159c002-2377-4fed-be8e-b14d30fa0c93") // Random uuid
		res, err := cli.GetOneUserWithResponse(ctx, id)
		assert.Nil(t, err)
		assert.Equal(t, res.StatusCode(), http.StatusNotFound)
	})
}

func TestGetUsersByStatus(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cli := mustCreateClient(t)

	t.Run("Expect blocked users to be returned", func(t *testing.T) {
		res, err := cli.GetUsersByStatus(ctx, &client.GetUsersByStatusParams{Status: "blocked"})
		assert.Nil(t, err)

		userList := unMarshalUserList(t, res)

		for _, u := range *userList {
			assert.Equalf(t, *u.Status, "blocked", "user's status expected to be equal: blocked")
			assertUser(t, &u)
		}

		assert.Equal(t, res.StatusCode, http.StatusOK)
	})

	t.Run("Expect unblocked users to be returned", func(t *testing.T) {
		res, err := cli.GetUsersByStatus(ctx, &client.GetUsersByStatusParams{Status: "blocked"})
		assert.Nil(t, err)

		userList := unMarshalUserList(t, res)

		want := "blocked"
		assert.Nil(t, err)

		for _, u := range *userList {
			assert.Equalf(t, *u.Status, want, "user's status expected to be equal: %s", want)
			assertUser(t, &u)
		}

		assert.Equal(t, res.StatusCode, http.StatusOK)
	})
}

func unMarshalUserList(t *testing.T, res *http.Response) *client.UserList {
	content, err := io.ReadAll(res.Body)
	assert.Nil(t, err)

	var userList client.UserList
	err = json.Unmarshal(content, &userList)
	assert.Nil(t, err)

	return &userList
}

func assertUser(t *testing.T, u *client.User) {
	assert.NotEmpty(t, u.FirstName)
	assert.NotEmpty(t, u.LastName)
	assert.NotEmpty(t, u.Email)
	assert.NotEmpty(t, u.Status)
}

func mustCreateClient(t *testing.T) *client.ClientWithResponses {
	cli, err := client.NewClientWithResponses("http://localhost:3001")
	assert.Nil(t, err)

	return cli
}
