package handler

import (
	"log"
	"net/http"
	"strconv"
	"task_tracker/internal/service"

	"github.com/go-chi/chi/v5"
)

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := service.GetAllTasks()
	w.Write([]byte(tasks))
}

func GetSingleTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)

	if err != nil {
		log.Fatalln(err)
	}

	task := service.GetSingleTask(uint(id))

	w.Write([]byte(task))
}
