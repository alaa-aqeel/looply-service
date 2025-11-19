package domain_errors

import "fmt"

type DbErrorCode string

const (
	// Common
	ErrUnknown DbErrorCode = "UNKNOWN_ERROR"

	// DB specific domain errors
	ErrDBUnique       DbErrorCode = "UNIQUE_ERROR"
	ErrDBForeignKey   DbErrorCode = "FOREIGN_KEY_ERROR"
	ErrDBNotNull      DbErrorCode = "NOT_NULL_ERROR"
	ErrDBCheck        DbErrorCode = "CHECK_ERROR"
	ErrDBInvalidInput DbErrorCode = "INVALID_INPUT_ERROR"
)

type DatabaseError struct {
	Code    DbErrorCode
	Message string
	Field   string // Optional: which column caused the error
	Details string // Optional: raw details
}

func (e *DatabaseError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}
