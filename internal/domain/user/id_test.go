package user_test

import (
	"testing"
	"users-service-cqrs/internal/domain/user"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	id := user.NewID()
	assert.NotNil(t, id.String())
}

func TestNewIDFromString(t *testing.T) {
	t.Run("Expect to return an error if provided string is not valid for an id", func(t *testing.T) {
		id, err := user.NewIDFromString("")
		assert.Nil(t, id)
		assert.NotNil(t, err)
	})

	t.Run("Expect to create ID from a string", func(t *testing.T) {
		id, err := user.NewIDFromString("4653e888-144c-471b-8be2-306a132df685")
		assert.Equal(t, id.String(), "4653e888-144c-471b-8be2-306a132df685")
		assert.Nil(t, err)
	})
}
