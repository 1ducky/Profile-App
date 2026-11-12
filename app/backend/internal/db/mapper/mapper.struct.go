package mapper

import (
	"errors"
)

type ErrorMapping struct {
	NotFound        error
	Duplicate       error
	RelatedNotFound error
	StillReferenced error
	RequiredField   error
	ValueTooLong    error
	Retryable       error
	Connection      error
}

var (
	ErrNoRows          = errors.New("Row not found")
	ErrDuplicate       = errors.New("Duplicate record")
	ErrRelatedNotFound = errors.New("Related record not found")
	ErrStillReferenced = errors.New("Still referenced")
	ErrRequiredField   = errors.New("Required field is empty")
	ErrValueTooLong    = errors.New("Value is too long")
	ErrRetryable       = errors.New("Retryable error")
	ErrConnection      = errors.New("Connection error")
)

func NewPersistanceErrMapper() *ErrorMapping {
	return &ErrorMapping{
		NotFound:        ErrNoRows,
		Duplicate:       ErrDuplicate,
		RelatedNotFound: ErrRelatedNotFound,
		StillReferenced: ErrStillReferenced,
		RequiredField:   ErrRequiredField,
		ValueTooLong:    ErrValueTooLong,
		Retryable:       ErrRetryable,
		Connection:      ErrConnection,
	}
}
