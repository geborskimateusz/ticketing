package customerr

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type SerializedError struct {
	Errors []string `json:"errors"`
}

type IApiError interface {
	Error() string
	SerializeErrors() []string
	GetStatusCode() int
}

type ApiError struct {
	IApiError
	StatusCode int
	Reason     string
}

type DatabaseConnectionError struct {
	*ApiError
}

func NewDataBaseConnectionError() DatabaseConnectionError {
	dbConnError := DatabaseConnectionError{
		&ApiError{
			StatusCode: http.StatusInternalServerError,
			Reason:     "Error connecting to database",
		},
	}
	return dbConnError
}

func (e DatabaseConnectionError) Error() string {
	return fmt.Sprintf(e.Reason)
}

func (e *DatabaseConnectionError) SerializeErrors() []string {
	return []string{e.Reason}
}

func (e *DatabaseConnectionError) GetStatusCode() int {
	return e.StatusCode
}

type RequestValidationError struct {
	*ApiError
}

func NewRequestValidationError(errors []validator.FieldError) RequestValidationError {
	reqValidationErr := RequestValidationError{
		&ApiError{
			StatusCode: http.StatusInternalServerError,
			Reason:     FiledErrorsAsString(errors),
		},
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
