package http

import (
	"errors"
	"net/http"
	"users-service-cqrs/internal/app"
	"users-service-cqrs/internal/app/command"
	"users-service-cqrs/internal/app/query"
	"users-service-cqrs/internal/domain/user"

	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/render"
)

type Handlers struct {
	application *app.App
}

func NewHandlers(application *app.App) *Handlers {
	return &Handlers{application: application}
}

var _ ServerInterface = (*Handlers)(nil)

func (h *Handlers) GetOneUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	userId, err := user.NewIDFromString(id.String())
	if err != nil {
		reply(w, r, NewErrorResponse(err, err.Error(), http.StatusBadRequest))
		return
	}

	u, err := h.application.Queries.OneUser.Handle(r.Context(), &userId)

	if errors.Is(err, user.ErrUserNotFound) {
		reply(w, r, NewErrorResponse(err, err.Error(), http.StatusNotFound))
		return
	}

	if err != nil {
		reply(w, r, NewErrorResponse(err, "An error occurred", http.StatusInternalServerError))
		return
	}

	render.Respond(w, r, UserResponse(u))
}

func (h *Handlers) BlockUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	userId, err := user.NewIDFromString(id.String())
	if err != nil {
		reply(w, r, NewErrorResponse(err, err.Error(), http.StatusBadRequest))
		return
	}

	err = h.application.Commands.BlockUser.Handle(r.Context(), command.BlockUser{UserId: &userId})

	if errors.Is(err, user.ErrUserNotFound) {
		reply(w, r, NewErrorResponse(err, err.Error(), http.StatusNotFound))
		return
	}

	render.Respond(w, r, NewGenericResponse("user blocked with success", http.StatusNoContent))
}

func (h *Handlers) UnblockUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	userId, err := user.NewIDFromString(id.String())
	if err != nil {
		reply(w, r, NewErrorResponse(err, err.Error(), http.StatusBadRequest))
		return
	}

	err = h.application.Commands.UnBlockUser.Handle(r.Context(), command.UnBlockUser{UserId: &userId})

	if errors.Is(err, user.ErrUserNotFound) {
		reply(w, r, NewErrorResponse(err, err.Error(), http.StatusNotFound))
		return
	}

	render.Respond(w, r, NewGenericResponse("user unblocked with success", http.StatusNoContent))
}

func (h *Handlers) GetUsersByStatus(w http.ResponseWriter, r *http.Request, params GetUsersByStatusParams) {
	users, err := h.application.Queries.AllUsers.Handle(r.Context(), query.AllUsers{Status: string(params.Status)})
	if err != nil {
		reply(w, r, NewErrorResponse(err, "an unexpected error occurred", http.StatusInternalServerError))
		return
	}

	render.Respond(w, r, UserListResponse(users))
}
