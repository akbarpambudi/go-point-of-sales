package category

import "context"

type (
	Repository interface {
		Create(ctx context.Context, entity *Category) error
		Update(ctx context.Context, id string, entity *Category) error
		Load(ctx context.Context, id string) (entity *Category, err error)
	}
)
