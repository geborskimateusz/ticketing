package errors

import "fmt"

type RequestValidationError struct {
	err string
}

func (e *RequestValidationError) Error() string {
	return fmt.Sprintf(e.err)
}
