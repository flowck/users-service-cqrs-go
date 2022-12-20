package http

import (
	"log"
	"net/http"
	"users-service-cqrs/internal/app/query"

	"github.com/go-chi/render"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
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

type ErrorResponse struct {
	Err     error  `json:"-"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	log.Printf("%s\n", e.Err)
	w.WriteHeader(e.Status)
	return nil
}

func NewErrorResponse(err error, msg string, status int) *ErrorResponse {
	return &ErrorResponse{
		Err:     err,
		Message: msg,
		Status:  status,
	}
}

func reply(w http.ResponseWriter, r *http.Request, payload render.Renderer) {
	if err := render.Render(w, r, payload); err != nil {
		render.Respond(w, r, NewErrorResponse(err, "Something unexpected happened", http.StatusInternalServerError))
	}
}

func NewGenericResponse(msg string, status int) *GenericResponse {
	return &GenericResponse{
		Message: &msg,
		Status:  &status,
	}
}
