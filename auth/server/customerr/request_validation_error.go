package customerr

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type RequestValidationError struct {
	statusCode int
	reason     string
}

func NewRequestValidationError(errors []validator.FieldError) DatabaseConnectionError {
	dbConnError := DatabaseConnectionError{}
	dbConnError.statusCode = http.StatusBadRequest
	dbConnError.reason = FiledErrorsAsString(errors)
	return dbConnError
}

func (e RequestValidationError) Error() string {
	return fmt.Sprintf(e.reason)
}

func (e *RequestValidationError) SerializeErrors() []string {
	return strings.Split(e.reason, Separator)
}
