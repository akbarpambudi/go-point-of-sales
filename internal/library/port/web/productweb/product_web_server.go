package productweb

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

func (p Server) CreateProduct(ctx echo.Context) error {
	var reqBody CreateProductJSONBody
	err := ctx.Bind(&reqBody)
	if err != nil {
		return err
	}

	var productVariants []command.ProductVariantDTO
	if reqBody.Variants != nil {
		for _, v := range *reqBody.Variants {
			productVariant := command.ProductVariantDTO{
				ID:    ptrval.StringVal(v.Id),
				Code:  ptrval.StringVal(v.Code),
				Name:  ptrval.StringVal(v.Name),
				Price: float64(ptrval.Float32Val(v.Price)),
			}
			productVariants = append(productVariants, productVariant)
		}
	}

	cmd := command.CreateProduct{
		ID:          ptrval.StringVal(reqBody.Id),
		Name:        ptrval.StringVal(reqBody.Name),
		CategoryRef: ptrval.StringVal(reqBody.CategoryRef),
		Variants:    productVariants,
	}

	err = p.application.Commands.CreateProduct.Handle(ctx.Request().Context(), cmd)
	if err != nil {
		return httphelper.WrapError(err)
	}
	return ctx.JSON(http.StatusCreated, nil)
}

func (p Server) GetProductById(ctx echo.Context, productId ProductIdParameter) error {
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

func (p Server) GetAllProducts(ctx echo.Context) error {
	products, err := p.application.Queries.GetAllProducts.Handle(ctx.Request().Context())
	if err != nil {
		return httphelper.WrapError(err)
	}
	var responseBody GetAllProductsSuccessResponse
	for _, p := range products {
		var variants []Variant

		for _, v := range p.Variants {
			variantPrice := float32(v.Price)
			variants = append(variants, Variant{
				Code:  &v.Code,
				Id:    &v.ID,
				Name:  &v.Name,
				Price: &variantPrice,
			})
		}

		responseBody = append(responseBody, Product{
			CategoryRef: &p.Category,
			Id:          &p.ID,
			Name:        &p.Name,
			Variants:    &variants,
		})
	}

	encodingErr := ctx.JSON(http.StatusOK, responseBody)

	if encodingErr != nil {
		return httphelper.WrapError(encodingErr)
	}

	return nil
}
