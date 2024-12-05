package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code    int      `json:"code"`
	Message string   `json:"msg"`
	Trace   []string `json:"trace"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码%d已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{Code: code, Message: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d,错误信息：%s", e.Code, e.Message)
}

func (e *Error) Details() []string {
	return e.Trace
}
func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.Trace = []string{}
	for _, d := range details {
		newError.Trace = append(newError.Trace, d)
	}
	return &newError
}

func (e *Error) StatusCode() int {
	switch e.Code {
	case Success.Code:
		return http.StatusOK
	case ServerError.Code:
		return http.StatusInternalServerError
	case InvalidParams.Code:
		return http.StatusOK
	case BusinessError.Code:
		return http.StatusOK
	case UnauthorizedAuthNotExist.Code:
		return http.StatusOK
	case UnauthorizedTokenError.Code:
		return http.StatusOK
	case UnauthorizedTokenGenerate.Code:
		return http.StatusOK
	case UnauthorizedTokenTimeout.Code:
		return http.StatusOK
	case TooManyRequests.Code:
		return http.StatusTooManyRequests
	}
	return http.StatusOK
}
