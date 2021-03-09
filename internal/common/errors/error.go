package errors

import "fmt"

type ErrorType = string

type POSError struct {
	errType ErrorType
	key     string
	message string
}

func NewPOSError(errType ErrorType, key string, message string) *POSError {
	return &POSError{errType: errType, key: key, message: message}
}

func (p POSError) ErrType() ErrorType {
	return p.errType
}

func (p POSError) Message() string {
	return p.message
}

func (p POSError) Key() string {
	return p.key
}

func (p POSError) Error() string {
	return fmt.Sprintf("%s: %s - %s", p.errType, p.key, p.message)
}
