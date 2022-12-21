package component

import (
	"context"
	"testing"
	"users-service-cqrs/tests/client"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBlockAndUnBlockUser(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cli := mustSetup(t)

	id, err := uuid.Parse("7e79baa9-318c-4286-9943-dd862eb65c1d")
	assert.Nil(t, err)

	_, err = cli.BlockUser(ctx, id)
	assert.Nil(t, err)

	res, err := cli.GetOneUserWithResponse(ctx, id)
	assert.Equalf(t, *res.JSON200.Status, "blocked", "expect user %s to be blocked", id.String())
	assert.Nil(t, err)

	_, err = cli.UnblockUser(ctx, id)
	assert.Nil(t, err)

	res, err = cli.GetOneUserWithResponse(ctx, id)
	assert.Equalf(t, *res.JSON200.Status, "unblocked", "expect user %s to be unblocked", id.String())
	assert.Nil(t, err)
}

func mustSetup(t *testing.T) *client.ClientWithResponses {
	cli, err := client.NewClientWithResponses("http://localhost:3001")
	assert.Nil(t, err)

	return cli
}
