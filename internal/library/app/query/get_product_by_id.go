package query

import "context"

type (
	GetProductByIDHandler interface {
		Handle(ctx context.Context, id string) (Product, error)
	}

	ProductReadModelProjector interface {
		LoadById(ctx context.Context, id string) (Product, error)
	}
)

type GetProductByIDHandlerImpl struct {
	projector ProductReadModelProjector
}

func NewGetProductByIDHandlerImpl(projector ProductReadModelProjector) *GetProductByIDHandlerImpl {
	return &GetProductByIDHandlerImpl{projector: projector}
}

func (g GetProductByIDHandlerImpl) Handle(ctx context.Context, id string) (Product, error) {
	return g.projector.LoadById(ctx, id)
}
