package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"task_tracker/internal/model"
	"task_tracker/pkg/apperr"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

type WriteResponse struct{}

func (r WriteResponse) SendData(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set(ContentType, ApplicationJson)

	responseData := model.Response[any]{Data: data, Status: statusCode}

	jsonResponse, err := json.Marshal(responseData)

	if err != nil {
		writeInternalError(w)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

func (r WriteResponse) SendAppErr(w http.ResponseWriter, appErr *apperr.AppErr) {
	w.Header().Set(ContentType, ApplicationJson)

	jsonResponse, err := json.Marshal(&appErr)

	if err != nil {
		writeInternalError(w)
		return
	}

	w.WriteHeader(appErr.Status)
	w.Write([]byte(jsonResponse))
}

func (r WriteResponse) SendErr(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set(ContentType, ApplicationJson)

	responseErr := model.Response[any]{Error: err.Error(), Status: statusCode}

	jsonResponse, err := json.Marshal(responseErr)

	if err != nil {
		writeInternalError(w)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

func writeInternalError(w http.ResponseWriter) {
	errJson := fmt.Sprintf(
		`{"status":%d,"error":"%s"}`,
		http.StatusInternalServerError,
		"Internal Server Error",
	)

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(errJson))
}
