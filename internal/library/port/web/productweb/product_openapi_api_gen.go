// Package productweb provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package productweb

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /product)
	GetAllProducts(ctx echo.Context) error

	// (POST /product)
	CreateProduct(ctx echo.Context) error

	// (GET /product/{productId})
	GetProductById(ctx echo.Context, productId ProductIdParameter) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetAllProducts converts echo context to params.
func (w *ServerInterfaceWrapper) GetAllProducts(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAllProducts(ctx)
	return err
}

// CreateProduct converts echo context to params.
func (w *ServerInterfaceWrapper) CreateProduct(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateProduct(ctx)
	return err
}

// GetProductById converts echo context to params.
func (w *ServerInterfaceWrapper) GetProductById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "productId" -------------
	var productId ProductIdParameter

	err = runtime.BindStyledParameter("simple", false, "productId", ctx.Param("productId"), &productId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter productId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetProductById(ctx, productId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/product", wrapper.GetAllProducts)
	router.POST(baseURL+"/product", wrapper.CreateProduct)
	router.GET(baseURL+"/product/:productId", wrapper.GetProductById)

}
