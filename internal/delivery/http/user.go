package routes

import (
	"task_tracker/internal/handler"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router, h *handler.UserHandler) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Get("/{id}", h.GetById)
		r.Post("/", h.Create)
	})
}
