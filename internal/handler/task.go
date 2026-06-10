package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"task_tracker/internal/model"
	"task_tracker/internal/service"

	"github.com/go-chi/chi/v5"
)

var taskService service.TaskService

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := taskService.GetAllTasks()

	if err != nil {
		writeResponse.SendAppErr(w, appErr.ConvertToAppErr(err))
		return
	}

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
		return
	}

	writeResponse.SendData(w, task, http.StatusOK)
}

func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.CreateTask

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		fmt.Println("err", err)
		writeResponse.SendAppErr(w, appErr.ConvertToAppErr(err))
		return
	}

	err = taskService.AddTask(&task)

	if err != nil {
		writeResponse.SendAppErr(w, appErr.ConvertToAppErr(err))
		return
	}

	w.WriteHeader(200)
}
