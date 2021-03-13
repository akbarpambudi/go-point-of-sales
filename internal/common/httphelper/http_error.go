package httphelper

import (
	"github.com/akbarpambudi/go-point-of-sales/internal/common/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type POSHTTPError struct {
	ErrorType  string         `json:"errorType"`
	MessageKey string         `json:"messageKey"`
	Message    string         `json:"message"`
	Children   []POSHTTPError `json:"children,omitempty"`
}

func WrapError(err error) *echo.HTTPError {
	if posErr, ok := err.(*errors.POSError); ok {
		return mapPosErrorToHTTPError(posErr)
	}

	if posErr, ok := err.(*errors.POSMultiError); ok {
		return mapPosMultiErrorToHTTPError(posErr)
	}

	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

func mapPosMultiErrorToHTTPError(err *errors.POSMultiError) *echo.HTTPError {
	httpError := mapPOSMultiErrorToPOSHTTPError(err)
	return mapPOSHTTPErrorToEchoHTTPError(err.ErrType(), httpError)
}

func mapPosErrorToHTTPError(err *errors.POSError) *echo.HTTPError {
	httpError := mapPOSErrorToPOSHTTPError(err)
	return mapPOSHTTPErrorToEchoHTTPError(err.ErrType(), httpError)
}

func mapPOSHTTPErrorToEchoHTTPError(errType errors.ErrorType, httpError POSHTTPError) *echo.HTTPError {
	switch errType {
	case errors.ErrorTypeAuthorizationError:
		return echo.NewHTTPError(http.StatusUnauthorized, httpError)
	case errors.ErrorTypeAuthenticationError:
		return echo.NewHTTPError(http.StatusBadRequest, httpError)
	case errors.ErrorTypeIllegalInputError:
		return echo.NewHTTPError(http.StatusBadRequest, httpError)
	case errors.ErrorTypeResourceNotFoundError:
		return echo.NewHTTPError(http.StatusNotFound, httpError)

	}
	return echo.NewHTTPError(http.StatusInternalServerError, httpError)
}

func mapPOSErrorToPOSHTTPError(err *errors.POSError) POSHTTPError {
	httpError := POSHTTPError{
		ErrorType:  err.ErrType(),
		MessageKey: err.Key(),
		Message:    err.Message(),
		Children:   nil,
	}
	return httpError
}

func mapPOSMultiErrorToPOSHTTPError(err *errors.POSMultiError) POSHTTPError {

	var children []POSHTTPError

	for _, child := range err.Children() {
		switch v := child.(type) {
		case *errors.POSError:
			children = append(children, mapPOSErrorToPOSHTTPError(v))
		case *errors.POSMultiError:
			children = append(children, mapPOSMultiErrorToPOSHTTPError(v))
		default:
			children = append(children, POSHTTPError{
				ErrorType:  errors.ErrorTypeInternalError,
				MessageKey: "error.unexpected",
				Message:    "unexpected error",
				Children:   nil,
			})

		}
	}

	httpError := POSHTTPError{
		ErrorType:  err.ErrType(),
		MessageKey: err.Key(),
		Message:    err.Message(),
		Children:   children,
	}
	return httpError
}
