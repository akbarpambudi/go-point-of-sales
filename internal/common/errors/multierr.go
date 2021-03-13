package errors

import (
	"bytes"
	"fmt"
	"go.uber.org/multierr"
)

type POSMultiError struct {
	*POSError
	children []error
}

func WrapMultiErr(errorType string, messageKey string, message string, err error) error {
	errs := multierr.Errors(err)
	var multiErrors []error
	for _, e := range errs {
		multiErrors = append(multiErrors, e)
	}

	return &POSMultiError{
		POSError: NewPOSError(errorType, messageKey, message),
		children: multiErrors,
	}
}

func (e POSMultiError) Children() []error {
	return e.children
}

func (e *POSMultiError) Error() string {
	var buf bytes.Buffer
	for _, child := range e.children {
		buf.WriteString(child.Error())
		buf.WriteString(";")
	}
	return fmt.Sprintf("%s : %s - [%s]", e.message, e.key, buf.String())
}
