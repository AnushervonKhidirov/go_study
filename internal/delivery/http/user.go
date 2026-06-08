package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("all_users"))
		})

		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("single_user"))
		})
	})
}
