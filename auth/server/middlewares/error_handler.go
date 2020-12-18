package middlewares

import (
	"log"
	"net/http"

	"github.com/geborskimateusz/auth/server/validation"
	"github.com/gin-gonic/gin"
)

// ErrorHandler is a middleware error for global error handling
func ErrorHandler() gin.HandlerFunc {
	return errorHandlerT(gin.ErrorTypeAny)
}

func errorHandlerT(errType gin.ErrorType) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		detectedErrors := c.Errors.ByType(errType)

		if len(detectedErrors) > 0 {
			err := detectedErrors[0].Err
			switch e := err.(type) {
			case validation.ApiError:
				log.Printf("%T -> %s", err, e.SerializeErrors())
				c.JSON(e.StatusCode, &validation.SerializedError{Errors: e.SerializeErrors()})
			default:
				c.JSON(http.StatusInternalServerError, &validation.SerializedError{Errors: []string{err.Error()}})
			}

			c.Abort()
			return
		}

	}
}
