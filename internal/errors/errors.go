package errors

import (
	"github.com/json-iterator/go"
	"github.com/xiaomeng79/istio-micro/cinit"
	"net/http"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// Error implements the error interface.
type Error struct {
	Id     string `json:"id"`
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// New generates a custom error.
func New(detail string, code int32) error {
	return &Error{
		Id:     cinit.Config.Service.Name,
		Code:   code,
		Detail: detail,
		Status: http.StatusText(int(code)),
	}
}

// Parse tries to parse a JSON string into an error. If that
// fails, it will set the given string as the error detail.
func Parse(err string) *Error {
	e := new(Error)
	errr := json.Unmarshal([]byte(err), e)
	if errr != nil {
		e.Detail = err
	}
	return e
}

// BadRequest generates a 400 error.
func BadRequest(detail string) error {
	return &Error{
		Id:     cinit.Config.Service.Name,
		Code:   400,
		Detail: detail,
		Status: http.StatusText(400),
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(detail string) error {
	return &Error{
		Id:     cinit.Config.Service.Name,
		Code:   401,
		Detail: detail,
		Status: http.StatusText(401),
	}
}

// Forbidden generates a 403 error.
func Forbidden(detail string) error {
	return &Error{
		Id:     cinit.Config.Service.Name,
		Code:   403,
		Detail: detail,
		Status: http.StatusText(403),
	}
}

// NotFound generates a 404 error.
func NotFound(detail string) error {
	return &Error{
		Id:     cinit.Config.Service.Name,
		Code:   404,
		Detail: detail,
		Status: http.StatusText(404),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(detail string) error {
	return &Error{
		Id:     cinit.Config.Service.Name,
		Code:   500,
		Detail: detail,
		Status: http.StatusText(500),
	}
}

// Conflict generates a 409 error.
func Conflict(detail string) error {
	return &Error{
		Id:     cinit.Config.Service.Name,
		Code:   409,
		Detail: detail,
		Status: http.StatusText(409),
	}
}
