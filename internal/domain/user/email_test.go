package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	t.Run("Expect email to be created from a string", func(t *testing.T) {
		email, err := NewEmail("john.doe@gmail.com")
		assert.Equal(t, email.String(), "john.doe@gmail.com")
		assert.Nil(t, err)
	})

	t.Run("Expect an error if provided email is empty", func(t *testing.T) {
		email, err := NewEmail("")
		assert.NotNil(t, err, "An empty argument should result in an error")
		assert.Nil(t, email)
	})

	t.Run("Expect an error if email provided is invalid", func(t *testing.T) {
		email, err := NewEmail("john.doegmail.com")
		assert.NotNil(t, err, "Invalid email should result in an error")
		assert.Nil(t, email)
	})
}
