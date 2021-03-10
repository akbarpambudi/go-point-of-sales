package web

import (
	"github.com/akbarpambudi/go-point-of-sales/internal/common/httphelper"
	"github.com/akbarpambudi/go-point-of-sales/internal/common/ptrval"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/command"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ProductWebServer struct {
	application app.Application
}

func NewProductWebServer(application app.Application) *ProductWebServer {
	return &ProductWebServer{application: application}
}

func (p ProductWebServer) CreateProduct(ctx echo.Context) error {
	var reqBody CreateProductJSONBody
	err := ctx.Bind(&reqBody)
	if err != nil {
		return err
	}

	var productVariants []command.ProductVariantDTO

	for _, v := range *reqBody.Variants {
		productVariant := command.ProductVariantDTO{
			ID:    ptrval.StringVal(v.Id),
			Code:  ptrval.StringVal(v.Code),
			Name:  ptrval.StringVal(v.Name),
			Price: float64(ptrval.Float32Val(v.Price)),
		}
		productVariants = append(productVariants, productVariant)
	}

	cmd := command.CreateProduct{
		ID:          *reqBody.Id,
		Name:        *reqBody.Name,
		CategoryRef: *reqBody.CategoryRef,
		Variants:    productVariants,
	}

	err = p.application.Commands.CreateProduct.Handle(ctx.Request().Context(), cmd)
	if err != nil {
		return httphelper.WrapError(err)
	}
	return nil
}

func (p ProductWebServer) GetProductById(ctx echo.Context, productId ProductIdParameter) error {
	productReadModel, err := p.application.Queries.GetProductById.Handle(ctx.Request().Context(), string(productId))
	if err != nil {
		return httphelper.WrapError(err)
	}

	var variants []Variant

	for _, v := range productReadModel.Variants {
		variantPrice := float32(v.Price)
		variants = append(variants, Variant{
			Code:  &v.Code,
			Id:    &v.ID,
			Name:  &v.Name,
			Price: &variantPrice,
		})
	}

	err = ctx.JSON(http.StatusOK, Product{
		CategoryRef: &productReadModel.Category,
		Id:          &productReadModel.ID,
		Name:        &productReadModel.Name,
		Variants:    &variants,
	})

	if err != nil {
		return httphelper.WrapError(err)
	}

	return nil
}
