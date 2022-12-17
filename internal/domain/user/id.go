package user

import "github.com/google/uuid"

type ID struct {
	value uuid.UUID
}

func NewID() ID {
	return ID{value: uuid.New()}
}

func NewIDFromString(value string) (ID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return ID{}, err
	}

	return ID{value: id}, nil
}

func (i *ID) isZero() bool {
	return i.value.ID() == 0
}

func (i *ID) String() string {
	return i.value.String()
}
