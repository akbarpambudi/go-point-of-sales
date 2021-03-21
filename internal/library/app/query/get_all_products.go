package query

import "context"

type (
	GetAllProductsHandler interface {
		Handle(ctx context.Context) (Products, error)
	}

	ProductsReadModelQueryProjector interface {
		GetAllProducts(ctx context.Context) (Products, error)
	}

	GetAllProductsHandlerImpl struct {
		projector ProductsReadModelQueryProjector
	}
)

func NewGetAllProductsHandlerImpl(projector ProductsReadModelQueryProjector) *GetAllProductsHandlerImpl {
	return &GetAllProductsHandlerImpl{projector: projector}
}

func (g GetAllProductsHandlerImpl) Handle(ctx context.Context) (Products, error) {
	return g.projector.GetAllProducts(ctx)
}
