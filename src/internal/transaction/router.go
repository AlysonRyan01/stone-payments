package transaction

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	handler *Handler
}

func NewRouter(h *Handler) *Router {
	return &Router{
		handler: h,
	}
}

func (rt *Router) Register(r chi.Router) {
	r.Post("/", rt.handler.Create)
}