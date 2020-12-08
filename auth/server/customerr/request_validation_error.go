package customerr

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type RequestValidationError struct {
	StatusCode int
	Reason     string
}

func NewRequestValidationError(errors []validator.FieldError) RequestValidationError {
	reqValidationErr := RequestValidationError{
		StatusCode: http.StatusBadRequest,
		Reason:     FiledErrorsAsString(errors),
	}
	return reqValidationErr
}

func (e RequestValidationError) Error() string {
	return fmt.Sprintf(e.Reason)
}

func (e *RequestValidationError) SerializeErrors() []string {
	return strings.Split(e.Reason, Separator)
}

func (e *RequestValidationError) GetStatusCode() int {
	return e.StatusCode
}
