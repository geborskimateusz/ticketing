package middlewares

import (
	"log"
	"net/http"

	"github.com/geborskimateusz/auth/server/customerr"
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
			case customerr.ApiError:
				log.Printf("%T -> %s", err, e.SerializeErrors())
				c.JSON(e.StatusCode, &customerr.SerializedError{Errors: e.SerializeErrors()})
			default:
				c.JSON(http.StatusInternalServerError, &customerr.SerializedError{Errors: []string{err.Error()}})
			}

			c.Abort()
			return
		}

	}
}
