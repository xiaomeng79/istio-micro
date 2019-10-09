package errors

import (
	"net/http"

	"github.com/xiaomeng79/istio-micro/cinit"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

//  Error implements the error interface.
type Error struct {
	ID     string `json:"id"`
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

func (e *Error) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

//  New generates a custom error.
func New(detail string, code int32) error {
	return &Error{
		ID:     cinit.Config.Service.Name,
		Code:   code,
		Detail: detail,
		Status: http.StatusText(int(code)),
	}
}

//  Parse tries to parse a JSON string into an error. If that//  fails, it will set the given string as the error detail.
func Parse(err string) *Error {
	e := new(Error)
	errr := json.Unmarshal([]byte(err), e)
	if errr != nil {
		e.Detail = err
	}
	return e
}

//  BadRequest generates a 400 error.
func BadRequest(detail string) error {
	return &Error{
		ID:     cinit.Config.Service.Name,
		Code:   http.StatusNotFound,
		Detail: detail,
		Status: http.StatusText(http.StatusNotFound),
	}
}

//  Unauthorized generates a 401 error.
func Unauthorized(detail string) error {
	return &Error{
		ID:     cinit.Config.Service.Name,
		Code:   http.StatusUnauthorized,
		Detail: detail,
		Status: http.StatusText(http.StatusUnauthorized),
	}
}

//  Forbidden generates a 403 error.
func Forbidden(detail string) error {
	return &Error{
		ID:     cinit.Config.Service.Name,
		Code:   http.StatusForbidden,
		Detail: detail,
		Status: http.StatusText(http.StatusForbidden),
	}
}

//  NotFound generates a 404 error.
func NotFound(detail string) error {
	return &Error{
		ID:     cinit.Config.Service.Name,
		Code:   http.StatusNotFound,
		Detail: detail,
		Status: http.StatusText(http.StatusNotFound),
	}
}

//  InternalServerError generates a 500 error.
func InternalServerError(detail string) error {
	return &Error{
		ID:     cinit.Config.Service.Name,
		Code:   http.StatusInternalServerError,
		Detail: detail,
		Status: http.StatusText(http.StatusInternalServerError),
	}
}

//  Conflict generates a 409 error.
func Conflict(detail string) error {
	return &Error{
		ID:     cinit.Config.Service.Name,
		Code:   http.StatusConflict,
		Detail: detail,
		Status: http.StatusText(http.StatusConflict),
	}
}
