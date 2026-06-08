package routes

import (
	"task_tracker/internal/handler"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handler.GetAllUsersHandler)
		r.Get("/{id}", handler.GetSingleUserHandler)
	})
}
