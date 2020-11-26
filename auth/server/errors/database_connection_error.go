package errors

import "fmt"

type DatabaseConnectionError struct{}

func (e *DatabaseConnectionError) Error() string {
	return fmt.Sprintf("Error connecting to database")
}
