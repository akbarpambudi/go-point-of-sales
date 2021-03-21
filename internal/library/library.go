package library

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/port/web/categoryweb"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/port/web/productweb"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type WebServiceOptions struct {
	Client *ent.Client
}

type WebServiceOptionsSetter func(options *WebServiceOptions)

func (o *WebServiceOptions) ApplySetters(optionsSetters ...WebServiceOptionsSetter) {
	for _, setTo := range optionsSetters {
		setTo(o)
	}
}

func NewWebService(ctx context.Context, optionsSetters ...WebServiceOptionsSetter) (http.Handler, func(), error) {
	e := echo.New()
	opts := WebServiceOptions{}
	opts.ApplySetters(optionsSetters...)
	e.Use(middleware.CORS())
	application, cleansingFunc, err := service.NewApplication(ctx, service.ApplicationOptions{
		Client: opts.Client,
	})
	if err != nil {
		return nil, nil, err
	}
	apiGroup := e.Group("/api")
	productweb.RegisterHandlers(apiGroup, productweb.NewServer(application))
	categoryweb.RegisterHandlers(apiGroup, categoryweb.NewServer(application))
	return e, cleansingFunc, nil
}
