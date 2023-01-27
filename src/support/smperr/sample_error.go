package smperr

import (
	"errors"
	"fmt"
	"net/http"
)

type appErr struct {
	error
	code int
	msg  string
}

type AppError interface {
	Code() int
	Msg() string
	Error() string
}

type BadRequestErr struct {
	*appErr
}

func (e *appErr) Code() int {
	return e.code
}
func (e *appErr) Msg() string {
	return e.msg
}

func BadRequestf(format string, msg ...any) *BadRequestErr {
	message := fmt.Sprintf(format, msg...)
	err := errors.New(message)
	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			error: err,
			msg:   message,
		},
	}
}
func BadRequestWrapf(err2 error, format string, msg ...any) *BadRequestErr {
	message := fmt.Sprintf(format, msg...)
	err := errors.New(message)
	err = errors.Join(err, err2)
	return &BadRequestErr{
		&appErr{
			code:  http.StatusBadRequest,
			error: err,
			msg:   message,
		},
	}
}
