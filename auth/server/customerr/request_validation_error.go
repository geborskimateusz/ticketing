package customerr

import (
	"github.com/go-playground/validator/v10"
)

type RequestValidationError struct {
	Errors []validator.FieldError
}

func (validationError RequestValidationError) Error() string {
	return FiledErrorsAsString(validationError.Errors)
}
