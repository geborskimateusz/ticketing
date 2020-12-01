package customerr

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type RequestValidationError struct {
	Errors []validator.FieldError
}

func (validationError RequestValidationError) Error() string {
	var sb strings.Builder

	for _, ve := range validationError.Errors {
		sb.WriteString("validation failed on field '" + ve.Field() + "'")
		sb.WriteString(", condition: " + ve.ActualTag())

		// Print condition parameters, e.g. oneof=red blue -> { red blue }
		if ve.Param() != "" {
			sb.WriteString(" { " + ve.Param() + " }")
		}

		if ve.Value() != nil && ve.Value() != "" {
			sb.WriteString(fmt.Sprintf(", actual: %v", ve.Value()))
		}
	}

	return sb.String()
}
