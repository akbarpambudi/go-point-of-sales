package product

import "context"

type (
	Repository interface {
		Create(ctx context.Context, entity *Product) error
		Update(ctx context.Context, id string, entity *Product) error
		Load(ctx context.Context, id string) (entity *Product, err error)
	}
)
