package apperr

import (
	"errors"
	"net/http"
)

type AppErr struct {
	Status  int     `json:"status"`
	Err     string  `json:"error"`
	Message *string `json:"message,omitempty"`
}

const (
	BadRequest          = "Bad Request"
	Unauthorized        = "Unauthorized"
	Forbidden           = "Forbidden"
	NotFound            = "Not Found"
	MethodNotAllowed    = "Method Not Allowed"
	Conflict            = "Conflict"
	InternalServerError = "Internal Server Error"
)

func (e AppErr) Error() string {
	return e.Err
}

func CreateErr(statusCode int, err string, message *string) *AppErr {
	return &(AppErr{Status: statusCode, Err: err, Message: message})
}

func ConvertToAppErr(err error) *AppErr {
	var appErr *AppErr

	if errors.As(err, &appErr) {
		return appErr
	}

	message := err.Error()
	return CreateErr(http.StatusInternalServerError, InternalServerError, &message)
}

// NOTE: Client error responses: 4XX
func BadRequestErr(message *string) *AppErr {
	return CreateErr(http.StatusBadRequest, BadRequest, message)
}

func UnauthorizedErr(message *string) *AppErr {
	return CreateErr(http.StatusUnauthorized, Unauthorized, message)
}

func ForbiddenErr(message *string) *AppErr {
	return CreateErr(http.StatusForbidden, Forbidden, message)
}

func NotFoundErr(message *string) *AppErr {
	return CreateErr(http.StatusNotFound, NotFound, message)
}

func MethodNotAllowedErr(message *string) *AppErr {
	return CreateErr(http.StatusMethodNotAllowed, MethodNotAllowed, message)
}

func ConflictErr(message *string) *AppErr {
	return CreateErr(http.StatusConflict, Conflict, message)
}

// NOTE: Server error responses: 5XX
func InternalServerErrorErr(message *string) *AppErr {
	return CreateErr(http.StatusInternalServerError, InternalServerError, message)
}
