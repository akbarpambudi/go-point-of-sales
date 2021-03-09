package httphelper

import (
	"github.com/akbarpambudi/go-point-of-sales/internal/common/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func WrapError(err error) *echo.HTTPError {
	if posErr, ok := err.(*errors.POSError); ok {
		return mapPosErrorToHTTPError(posErr)
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

func mapPosErrorToHTTPError(err *errors.POSError) *echo.HTTPError {
	switch err.ErrType() {
	case errors.ErrorTypeAuthorizationError:
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	case errors.ErrorTypeAuthenticationError:
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	case errors.ErrorTypeIllegalInputError:
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
