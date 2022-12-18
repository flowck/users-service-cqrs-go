package user

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewID(t *testing.T) {
	id := NewID()
	assert.NotNil(t, id.String())
}

func TestNewIDFromString(t *testing.T) {
	id, err := NewIDFromString("")
	assert.Equal(t, id.isZero(), true)
	assert.NotNil(t, err)

	id, err = NewIDFromString("4653e888-144c-471b-8be2-306a132df685")
	assert.Equal(t, id.String(), "4653e888-144c-471b-8be2-306a132df685")
	assert.Nil(t, err)
}
