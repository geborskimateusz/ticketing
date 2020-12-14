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

func (e ApiError) Error() string {
	return fmt.Sprintf(e.Reason)
}

func (e *ApiError) SerializeErrors() []string {
	return []string{e.Reason}
}

func (e *ApiError) GetStatusCode() int {
	return e.StatusCode

}

type DatabaseConnectionError struct {
	*ApiError
}

func NewDataBaseConnectionError() ApiError {
	dbConnError := &DatabaseConnectionError{
		&ApiError{
			StatusCode: http.StatusInternalServerError,
			Reason:     "Error connecting to database",
		},
	}
	return *dbConnError.ApiError
}

type RequestValidationError struct {
	*ApiError
}

func NewRequestValidationError(errors []validator.FieldError) ApiError {
	reqValidationErr := &RequestValidationError{
		&ApiError{
			StatusCode: http.StatusBadRequest,
			Reason:     FiledErrorsAsString(errors),
		},
	}
	return *reqValidationErr.ApiError
}

func (e *RequestValidationError) SerializeErrors() []string {
	return strings.Split(e.Reason, Separator)
}

type NotFoundError struct {
	*ApiError
}

func NewNotFoundError() ApiError {
	notFoundErr := &NotFoundError{
		&ApiError{
			StatusCode: http.StatusNotFound,
			Reason:     "Not found",
		},
	}
	return *notFoundErr.ApiError
}