package customerr

import (
	"fmt"
	"net/http"
)

type DatabaseConnectionError struct {
	statusCode int
	reason     string
}

func NewDataBaseConnectionError() DatabaseConnectionError {
	dbConnError := DatabaseConnectionError{}
	dbConnError.statusCode = http.StatusInternalServerError
	dbConnError.reason = "Error connecting to database"
	return dbConnError
}

func (e *DatabaseConnectionError) Error() string {
	return fmt.Sprintf(e.reason)
}

func (e *DatabaseConnectionError) SerializeErrors() []string {
	return []string{e.reason}
}
