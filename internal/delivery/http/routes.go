package routes

import (
	"net/http"
	"task_tracker/internal/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRoutes(handlers *handler.Handler) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Route("/api/v1", func(r chi.Router) {
		UserRoutes(r, handlers.User)
		TaskRoutes(r, handlers.Task)
	})

	return r
}
