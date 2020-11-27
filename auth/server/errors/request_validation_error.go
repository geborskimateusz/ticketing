package errors

import "fmt"

type RequestValidationError struct {
	// Err string
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *RequestValidationError) Error() string {
	// return fmt.Sprintf(e.Err)
	return fmt.Sprintf(e.Message)
}
