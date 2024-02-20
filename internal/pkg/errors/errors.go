package errors

import (
	"fmt"
	"runtime"
)

type Errors struct {
	Code int
	Message string
	callstack []uintptr
}

func New(code int, message string) error {
	return &Errors{
		Code: code,
		Message: message,
		callstack: initStackTrace(),
	}
}

func (e *Errors) Error() string {
	return fmt.Sprintf("Error code: %d, Message: %s", e.Code, e.Message)
}

func (e *Errors) StackTraceError() string {
	return ""
}

func initStackTrace() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(2, pcs[:])

	return pcs[:n]
}
