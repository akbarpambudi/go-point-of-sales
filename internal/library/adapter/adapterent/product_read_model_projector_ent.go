package adapterent

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/query"
)

func (p ProductRepository) LoadById(ctx context.Context, id string) (query.Product, error) {
	domainEntity, err := p.Load(ctx, id)
	if err != nil {
		return query.Product{}, err
	}

	return query.Product{
		ID:       domainEntity.ID(),
		Name:     domainEntity.Name(),
		Category: domainEntity.CategoryRef(),
	}, nil
}
