package library

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/port/web/categoryweb"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/port/web/productweb"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewWebService(ctx context.Context) (http.Handler, func(), error) {
	e := echo.New()

	application, cleansingFunc, err := service.NewApplication(ctx)
	if err != nil {
		return nil, nil, err
	}
	apiGroup := e.Group("/api")
	productweb.RegisterHandlers(apiGroup, productweb.NewServer(application))
	categoryweb.RegisterHandlers(apiGroup, categoryweb.NewServer(application))
	return e, cleansingFunc, nil
}
