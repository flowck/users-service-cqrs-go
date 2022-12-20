package http

import (
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/go-chi/render"
	"net/http"
	"users-service-cqrs/internal/app"
	"users-service-cqrs/internal/app/query"
)

type Handlers struct {
	application *app.App
}

func NewHandlers(application *app.App) *Handlers {
	return &Handlers{application: application}
}

var _ ServerInterface = (*Handlers)(nil)

func (h *Handlers) BlockUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *Handlers) UnblockUser(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	//TODO implement me
	panic("implement me")
}

func (h *Handlers) GetUsersByStatus(w http.ResponseWriter, r *http.Request, params GetUsersByStatusParams) {
	users, err := h.application.Queries.AllBlockedUser.Handle(r.Context(), query.AllUsers{Status: string(params.Status)})
	if err != nil {
		return
	}

	render.Respond(w, r, UserListResponse(users))
}
