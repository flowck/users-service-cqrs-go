package adapters_test

import (
	"context"
	"testing"
	"users-service-cqrs/internal/adapters"
	"users-service-cqrs/internal/app/query"
	"users-service-cqrs/internal/common/config"
	"users-service-cqrs/internal/common/psql"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestPsqlReadUserRepo_FindAll(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cfg := config.New()
	db := psql.Connect(cfg.PsqlUri)

	repo := adapters.NewPsqlReadUserRepo(db)
	userList, err := repo.FindAll(ctx, query.AllUsers{Status: "blocked"})

	assert.Nil(t, err)
	assert.NotEmpty(t, userList, "Should return a list of users")

	for _, u := range userList {
		assert.Equal(t, u.Status, "blocked")
	}
}
