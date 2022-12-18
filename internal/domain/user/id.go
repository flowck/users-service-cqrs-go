package user

import (
	"errors"

	"github.com/google/uuid"
)

type ID struct {
	value uuid.UUID
}

func NewID() ID {
	return ID{value: uuid.New()}
}

func NewIDFromString(id string) (ID, error) {
	if id == "" {
		return ID{}, errors.New("id is empty")
	}

	newId, err := uuid.Parse(id)
	if err != nil {
		return ID{}, errors.New("the id is invalid")
	}

	return ID{value: newId}, nil
}

func (i *ID) isZero() bool {
	return i.value.ID() == 0
}

func (i *ID) String() string {
	return i.value.String()
}
