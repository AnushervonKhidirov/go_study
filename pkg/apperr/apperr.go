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

func (e AppErr) CreateErr(statusCode int, err string, message *string) *AppErr {
	e.Status = statusCode
	e.Err = err
	e.Message = message
	return &e
}

func (e AppErr) ConvertToAppErr(err error) *AppErr {
	var appErrType *AppErr

	if errors.As(err, &appErrType) {
		return appErrType
	}

	return e.CreateErr(http.StatusInternalServerError, InternalServerError, nil)
}

// NOTE: Client error responses: 4XX
func (e AppErr) BadRequestErr(message *string) *AppErr {
	return e.CreateErr(http.StatusBadRequest, BadRequest, message)
}

func (e AppErr) UnauthorizedErr(message *string) *AppErr {
	return e.CreateErr(http.StatusUnauthorized, Unauthorized, message)
}

func (e AppErr) ForbiddenErr(message *string) *AppErr {
	return e.CreateErr(http.StatusForbidden, Forbidden, message)
}

func (e AppErr) NotFoundErr(message *string) *AppErr {
	return e.CreateErr(http.StatusNotFound, NotFound, message)
}

func (e AppErr) MethodNotAllowedErr(message *string) *AppErr {
	return e.CreateErr(http.StatusMethodNotAllowed, MethodNotAllowed, message)
}

func (e AppErr) ConflictErr(message *string) *AppErr {
	return e.CreateErr(http.StatusConflict, Conflict, message)
}

// NOTE: Server error responses: 5XX
func (e AppErr) InternalServerErrorErr(message *string) *AppErr {
	return e.CreateErr(http.StatusInternalServerError, InternalServerError, message)
}
