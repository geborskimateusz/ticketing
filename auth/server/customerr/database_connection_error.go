package customerr

import (
	"fmt"
	"net/http"
)

type DatabaseConnectionError struct {
	StatusCode int
	Reason     string
}

func NewDataBaseConnectionError() DatabaseConnectionError {
	dbConnError := DatabaseConnectionError{
		StatusCode: http.StatusInternalServerError,
		Reason:     "Error connecting to database",
	}
	return dbConnError
}

func (e DatabaseConnectionError) Error() string {
	return fmt.Sprintf(e.Reason)
}

func (e *DatabaseConnectionError) SerializeErrors() []string {
	return []string{e.Reason}
}
