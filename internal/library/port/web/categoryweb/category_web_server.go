package categoryweb

import (
	"github.com/akbarpambudi/go-point-of-sales/internal/common/httphelper"
	"github.com/akbarpambudi/go-point-of-sales/internal/common/ptrval"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/command"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
	application app.Application
}

func NewServer(application app.Application) *Server {
	return &Server{application: application}
}

func (s Server) CreateCategory(ctx echo.Context) error {

	var reqBody CreateCategoryJSONRequestBody

	if err := ctx.Bind(&reqBody); err != nil {
		return httphelper.WrapError(err)
	}

	err := s.application.Commands.CreateCategory.Handle(ctx.Request().Context(), command.CreateCategory{
		ID:   ptrval.StringVal(reqBody.Id),
		Name: ptrval.StringVal(reqBody.Name),
	})

	if err != nil {
		return httphelper.WrapError(err)
	}

	return ctx.JSON(http.StatusCreated, nil)
}
