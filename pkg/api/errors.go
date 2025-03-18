package api

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"`
	HTTPStatusCode int   `json:"-"`

	Message   string `json:"message"`
	AppCode   int64  `json:"code,omitempty"`
	ErrorText string `json:"error,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var (
	ErrNotFound            = &ErrResponse{HTTPStatusCode: 404, Message: "Resource not found."}
	ErrBadRequest          = &ErrResponse{HTTPStatusCode: 400, Message: "Bad request"}
	ErrInternalServerError = &ErrResponse{HTTPStatusCode: 500, Message: "Internal Server Error"}
)

func NewError(err error, message string, statusCode int) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: statusCode,
		Message:        message,
		ErrorText:      err.Error(),
	}
}

func DuplicateKeys(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 409,
		Message:        "Duplicate Key",
		ErrorText:      err.Error(),
	}
}
