package response

import (
	"encoding/json"
	"net/http"
	"task_tracker/internal/model"
	"task_tracker/pkg/apperr"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

func SendData[Type any](w http.ResponseWriter, data Type, statusCode int) {
	w.Header().Set(ContentType, ApplicationJson)

	responseData := model.Response[Type]{Data: data, Status: statusCode}
	jsonResponse, err := json.Marshal(responseData)

	if err != nil {
		writeInternalError(w)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(jsonResponse)
}

func SendErr(w http.ResponseWriter, appErr *apperr.AppErr) {
	w.Header().Set(ContentType, ApplicationJson)

	jsonResponse, err := json.Marshal(&appErr)

	if err != nil {
		writeInternalError(w)
		return
	}

	w.WriteHeader(appErr.Status)
	w.Write([]byte(jsonResponse))
}

func writeInternalError(w http.ResponseWriter) {
	errResp := model.Response[any]{Status: http.StatusInternalServerError, Error: "Internal Server Error"}

	jsonResp, err := json.Marshal(errResp)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status":500,"error":"Internal Server Error"}`))
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonResp)
}
