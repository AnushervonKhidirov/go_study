package handler

import (
	"net/http"
	"task_tracker/internal/model"
	"task_tracker/internal/service"
	"task_tracker/pkg/apperr"
	"task_tracker/pkg/helpers"
	"task_tracker/pkg/response"

	"github.com/go-chi/chi/v5"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.service.GetAll(r.Context())

	if err != nil {
		response.SendErr(w, apperr.ConvertToAppErr(err))
		return
	}

	response.SendData(w, tasks, http.StatusOK)
}

func (h *TaskHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ParseID(chi.URLParam(r, "id"))

	if err != nil {
		response.SendErr(w, apperr.ConvertToAppErr(err))
		return
	}

	task, err := h.service.GetById(int(id), r.Context())

	if err != nil {
		response.SendErr(w, apperr.ConvertToAppErr(err))
		return
	}

	response.SendData(w, *task, http.StatusOK)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var t model.Task

	err := helpers.ValidateAndDecode(r.Body, &t)

	if err != nil {
		message := err.Error()
		response.SendErr(w, apperr.BadRequestErr(&message))
		return
	}

	err = h.service.Create(&t, r.Context())

	if err != nil {
		response.SendErr(w, apperr.ConvertToAppErr(err))
		return
	}

	response.SendData(w, &t, http.StatusCreated)
}
