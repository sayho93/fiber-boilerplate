package errors

import (
	"fmt"
	"github.com/gofiber/fiber/v2/utils"
	"runtime"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Stack   string `json:"stack"`
}

func New(code int, message string) *Error {
	stackTraceBuf := make([]byte, 1<<10) // 1kb
	runtime.Stack(stackTraceBuf, true)
	err :=
		&Error{
			Code:    code,
			Message: utils.StatusMessage(code),
			Stack:   string(stackTraceBuf),
		}
	if len(message) > 0 {
		err.Message = message
	}

	return err
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, message: %s, stack: %s", e.Code, e.Message, e.Stack)
}

func (e *Error) StackTrace() string {
	return e.Stack
}
