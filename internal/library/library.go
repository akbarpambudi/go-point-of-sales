package library

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/port/web"
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

	web.RegisterHandlers(e, web.NewProductWebServer(application))

	return e, cleansingFunc, nil
}
