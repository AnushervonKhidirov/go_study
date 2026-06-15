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

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll()

	if err != nil {
		response.SendAppErr(w, apperr.ConvertToAppErr(err))
		return
	}

	response.SendData(w, users, http.StatusOK)
}

func (h *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 0, 64)

	if err != nil {
		response.SendAppErr(w, apperr.ConvertToAppErr(err))
		return
	}

	user, err := h.service.GetById(int(id))

	if err != nil {
		response.SendAppErr(w, apperr.ConvertToAppErr(err))
		return
	}

	response.SendData(w, *user, http.StatusOK)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var u model.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		message := err.Error()
		response.SendAppErr(w, apperr.BadRequestErr(&message))
		return
	}

	err = validation.Validate(&u)

	if err != nil {
		message := err.Error()
		response.SendAppErr(w, apperr.BadRequestErr(&message))
		return
	}

	err = h.service.Create(&u)

	if err != nil {
		response.SendAppErr(w, apperr.ConvertToAppErr(err))
		return
	}

	response.SendData(w, &u, http.StatusCreated)
}
