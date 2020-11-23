package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type customError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *customError) Error() string {
	return err.Message
}

// Middleware error for global handler
func ErrorHandler() gin.HandlerFunc {
	return errorHandlerT(gin.ErrorTypeAny)
}

func errorHandlerT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errors := c.Errors.ByType(errType)

		log.Println("Handle APP error")
		if len(errors) > 0 {
			err := errors[0].Err
			var parsedError *customError
			switch err.(type) {
			case *customError:
				parsedError = err.(*customError)
			default:
				parsedError = &customError{
					Code:    http.StatusInternalServerError,
					Message: "Internal Server Error",
				}
			}

			c.IndentedJSON(parsedError.Code, parsedError)
			c.Abort()
			return
		}
	}
}
