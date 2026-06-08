package handler

import (
	"log"
	"net/http"
	"strconv"
	"task_tracker/internal/service"

	"github.com/go-chi/chi/v5"
)

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := service.GetAllUsers()
	w.Write([]byte(users))
}

func GetSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)

	if err != nil {
		log.Fatalln(err)
	}

	user := service.GetSingleUser(uint(id))
	w.Write([]byte(user))
}
