package middlewares

import (
	"fmt"
	"net/http"

	"github.com/geborskimateusz/auth/server/customerr"
	"github.com/gin-gonic/gin"
)

type customError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *customError) Error() string {
	return err.Message
}

// ErrorHandler is a middleware error for global error handling
func ErrorHandler() gin.HandlerFunc {
	return errorHandlerT(gin.ErrorTypeAny)
}

func errorHandlerT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errors := c.Errors.ByType(errType)

		if len(errors) > 0 {
			err := errors[0].Err

			switch err.(type) {
			case *customerr.RequestValidationError:
				fmt.Println("is request validation error")
			case *customerr.DatabaseConnectionError:
				fmt.Println("is database connection error")
			default:
				fmt.Println("none of them")
			}

			parsedError := &customError{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			}

			c.JSON(parsedError.Code, parsedError)
			c.Abort()
			return
		}
	}
}
