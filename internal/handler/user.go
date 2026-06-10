package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"task_tracker/internal/model"
	"task_tracker/internal/service"
	"task_tracker/pkg/validation"

	"github.com/go-chi/chi/v5"
)

var userService service.UserService

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := userService.GetAllUsers()

	if err != nil {
		writeResponse.SendErr(w, err, http.StatusInternalServerError)
		return
	}

	writeResponse.SendData(w, users, http.StatusOK)
}

func GetSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 0, 64)

	if err != nil {
		writeResponse.SendErr(w, err, http.StatusInternalServerError)
		return
	}

	user, err := userService.GetSingleUser(uint(id))

	if err != nil {
		writeResponse.SendAppErr(w, appErr.ConvertToAppErr(err))
		return
	}

	writeResponse.SendData(w, user, http.StatusOK)
}

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
	var u model.CreateUser

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		message := err.Error()
		writeResponse.SendAppErr(w, appErr.BadRequestErr(&message))
		return
	}

	err = validation.Validate(&u)

	if err != nil {
		message := err.Error()
		writeResponse.SendAppErr(w, appErr.BadRequestErr(&message))
		return
	}

	err = userService.AddUser(&u)

	if err != nil {
		writeResponse.SendAppErr(w, appErr.ConvertToAppErr(err))
		return
	}

	w.WriteHeader(200)
}
