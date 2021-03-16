package query

import "context"

type (
	GetCategoryByIDHandler interface {
		Handle(ctx context.Context, id string) (Category, error)
	}

	CategoryReadModelProjector interface {
		LoadById(ctx context.Context, ID string) (Category, error)
	}

	GetCategoryByIDHandlerImpl struct {
		projector CategoryReadModelProjector
	}
)

func NewGetCategoryByIDHandlerImpl(projector CategoryReadModelProjector) *GetCategoryByIDHandlerImpl {
	return &GetCategoryByIDHandlerImpl{projector: projector}
}

func (g *GetCategoryByIDHandlerImpl) Handle(ctx context.Context, id string) (Category, error) {
	return g.projector.LoadById(ctx, id)
}
