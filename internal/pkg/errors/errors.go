package errors

import (
	"fmt"
	"runtime"
)

type Errors struct {
	Code      int
	Message   string
	callstack []uintptr
}

func New(code int, message string) error {
	return &Errors{
		Code:      code,
		Message:   message,
		callstack: initStackTrace(),
	}
}

func (e *Errors) Error() string {
	return fmt.Sprintf("Error code: %d, Message: %s", e.Code, e.Message)
}

func (e *Errors) StackTraceError() string {
	var stackMessage string
	for _, pc := range e.callstack {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		stackMessage += fmt.Sprintf("%s:%d %s\n", file, line, fn.Name())
	}

	return stackMessage
}

func initStackTrace() []uintptr {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(2, pcs[:])

	return pcs[:n]
}
