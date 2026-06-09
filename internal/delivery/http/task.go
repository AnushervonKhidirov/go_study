package routes

import (
	"task_tracker/internal/handler"

	"github.com/go-chi/chi/v5"
)

func TaskRoutes(r chi.Router) {
	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", handler.GetAllTasksHandler)
		r.Get("/{id}", handler.GetSingleTaskHandler)
	})
}
