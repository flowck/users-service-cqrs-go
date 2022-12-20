package http

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	"users-service-cqrs/internal/app/query"
)

func UserResponse(u *query.User) *User {
	e := openapi_types.Email(u.Email)
	id, _ := uuid.Parse(u.Id)

	return &User{
		Id:        &id,
		Email:     &e,
		FirstName: &u.FirstName,
		LastName:  &u.LastName,
		Status:    &u.Status,
	}
}

func UserListResponse(list []*query.User) UserList {
	result := make(UserList, len(list))

	for idx, u := range list {
		result[idx] = *UserResponse(u)
	}

	return result
}
