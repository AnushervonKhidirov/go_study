package routes

import (
	"task_tracker/internal/handler"

	"github.com/go-chi/chi/v5"
)

func TaskRoutes(r chi.Router, h *handler.TaskHandler) {
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", h.GetAll)
		r.Get("/{id}", h.GetById)
		r.Post("/", h.Create)
	})
}
