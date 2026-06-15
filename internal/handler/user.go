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

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll()

	if err != nil {
		response.SendErr(w, apperr.ConvertToAppErr(err))
		return
	}

	response.SendData(w, users, http.StatusOK)
}

func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ParseID(chi.URLParam(r, "id"))

	if err != nil {
		response.SendErr(w, apperr.ConvertToAppErr(err))
		return
	}

	user, err := h.service.GetById(int(id))

	if err != nil {
		response.SendErr(w, apperr.ConvertToAppErr(err))
		return
	}

	response.SendData(w, *user, http.StatusOK)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u model.User

	err := helpers.ValidateAndDecode(r.Body, &u)

	if err != nil {
		message := err.Error()
		response.SendErr(w, apperr.BadRequestErr(&message))
		return
	}

	err = h.service.Create(&u)

	if err != nil {
		response.SendErr(w, apperr.ConvertToAppErr(err))
		return
	}

	response.SendData(w, &u, http.StatusCreated)
}
