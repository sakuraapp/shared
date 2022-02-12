package resource

import (
	"github.com/go-chi/render"
	"net/http"
)

type Response struct {
	Status int `json:"status"`
}

func (res *Response) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, res.Status)

	return nil
}

func NewResponse(status int) *Response {
	return &Response{Status: status}
}

type ErrResponse struct {
	Response
	Message *string `json:"message,omitempty"`
}

func (res *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, res.Status)

	return nil
}

func NewError(status int, message *string) *ErrResponse {
	return &ErrResponse{
		Response{status},
		message,
	}
}

func NewErrorMessage(message string) *string {
	return &message
}

var ErrBadRequest = NewError(http.StatusBadRequest, nil)
var ErrInternalError = NewError(http.StatusInternalServerError, nil)
var ErrUnauthorized = NewError(http.StatusUnauthorized, nil)
var ErrForbidden = NewError(http.StatusForbidden, nil)
var ErrNotFound = NewError(http.StatusNotFound, nil)