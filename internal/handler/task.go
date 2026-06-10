package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task_tracker/internal/model"
	"task_tracker/internal/service"
	"task_tracker/pkg/apperr"
	"task_tracker/pkg/response"
	"task_tracker/pkg/validation"

	"github.com/go-chi/chi/v5"
)

var taskService service.TaskService

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := taskService.GetAllTasks()

	if err != nil {
		response.SendAppErr(w, apperr.ConvertToAppErr(err))
		return
	}

	response.SendData(w, tasks, http.StatusOK)
}

func GetSingleTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)

	if err != nil {
		response.SendAppErr(w, apperr.ConvertToAppErr(err))
	}

	task, err := taskService.GetSingleTask(uint(id))

	if err != nil {
		response.SendAppErr(w, apperr.ConvertToAppErr(err))
		return
	}

	response.SendData(w, task, http.StatusOK)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t model.CreateTask

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		message := err.Error()
		response.SendAppErr(w, apperr.BadRequestErr(&message))
		return
	}

	err = validation.Validate(&t)

	if err != nil {
		message := err.Error()
		response.SendAppErr(w, apperr.BadRequestErr(&message))
		return
	}

	err = taskService.AddTask(&t)

	if err != nil {
		response.SendAppErr(w, apperr.ConvertToAppErr(err))
		return
	}

	w.WriteHeader(200)
}
