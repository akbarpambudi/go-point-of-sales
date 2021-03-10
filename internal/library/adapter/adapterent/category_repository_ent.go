package adapterent

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/category"
	"github.com/google/uuid"
)

type CategoryRepositoryEnt struct {
	client *ent.Client
}

func NewCategoryRepositoryEnt(client *ent.Client) *CategoryRepositoryEnt {
	return &CategoryRepositoryEnt{client: client}
}

func (c CategoryRepositoryEnt) Create(ctx context.Context, entity *category.Category) error {
	tx, err := c.client.Tx(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Category.
		Create().
		SetID(uuid.MustParse(entity.ID())).
		SetName(entity.Name()).
		Save(ctx)

	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

func (c CategoryRepositoryEnt) Update(ctx context.Context, id string, entity *category.Category) error {
	panic("implement me")
}

func (c CategoryRepositoryEnt) Load(ctx context.Context, id string) (entity *category.Category, err error) {
	panic("implement me")
}
