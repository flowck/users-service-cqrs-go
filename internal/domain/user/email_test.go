package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEmail(t *testing.T) {
	email, err := NewEmail("john.doe@gmail.com")
	assert.Equal(t, email.String(), "john.doe@gmail.com")
	assert.Nil(t, err)

	email, err = NewEmail("")
	assert.NotNil(t, err, "An empty argument should result in an error")

	email, err = NewEmail("john.doegmail.com")
	assert.NotNil(t, err, "Invalid email should result in an error")
}
