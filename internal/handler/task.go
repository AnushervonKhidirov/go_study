package handler

import (
	"net/http"
	"strconv"
	"task_tracker/internal/service"

	"github.com/go-chi/chi/v5"
)

var taskService service.TaskService

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := taskService.GetAllTasks()

	if err != nil {
		writeResponse.SendAppErr(w, appErr.ConvertToAppErr(err))
	}

	// w.Write([]byte(tasks))
	writeResponse.SendData(w, tasks, http.StatusOK)
}

func GetSingleTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)

	if err != nil {
		writeResponse.SendAppErr(w, appErr.ConvertToAppErr(err))
	}

	task, err := taskService.GetSingleTask(uint(id))

	if err != nil {
		writeResponse.SendAppErr(w, appErr.ConvertToAppErr(err))
	}

	writeResponse.SendData(w, task, http.StatusOK)
}
