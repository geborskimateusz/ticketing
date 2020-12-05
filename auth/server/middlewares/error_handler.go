package middlewares

import (
	"net/http"

	"github.com/geborskimateusz/auth/server/customerr"
	"github.com/gin-gonic/gin"
)

type customError struct {
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
			case *customerr.DatabaseConnectionError:
				c.JSON(err.StatusCode, gin.H{"errors": err.SerializeErrors()})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"errors": "Something went wrong."})
			}

			c.Abort()
			return
		}
	}
}
