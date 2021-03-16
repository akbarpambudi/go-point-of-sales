package adapterent

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/query"
	"github.com/google/uuid"
)

type CategoryReadModelProjectEnt struct {
	client *ent.Client
}

func NewCategoryReadModelProjectEnt(client *ent.Client) *CategoryReadModelProjectEnt {
	return &CategoryReadModelProjectEnt{client: client}
}

func (c CategoryReadModelProjectEnt) LoadById(ctx context.Context, id string) (query.Category, error) {
	model, err := c.client.Category.Get(ctx, uuid.MustParse(id))
	if err != nil {
		if ent.IsNotFound(err) {
			return query.Category{}, query.ErrCategoryResourceNotFound
		}
		return query.Category{}, err
	}

	return query.Category{
		ID:   model.ID.String(),
		Name: model.Name,
	}, nil
}
