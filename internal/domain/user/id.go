package user

import (
	"errors"
	"github.com/google/uuid"
)

type ID struct {
	value uuid.UUID
}

func NewID() *ID {
	return &ID{value: uuid.New()}
}

func NewIDFromString(value string) (*ID, error) {
	if value == "" {
		return nil, errors.New("id is empty")
	}

	id, err := uuid.Parse(value)
	if err != nil {
		return nil, errors.New("the id is invalid")
	}

	return &ID{value: id}, nil
}

func (i *ID) isZero() bool {
	return i.value.ID() == 0
}

func (i *ID) String() string {
	return i.value.String()
}
