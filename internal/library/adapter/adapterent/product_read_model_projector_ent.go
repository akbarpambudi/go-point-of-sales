package adapterent

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/query"
)

func (p ProductRepository) LoadById(ctx context.Context, id string) (query.Product, error) {
	domainEntity, err := p.Load(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return query.Product{}, query.ErrProductResourceNotFound
		}
		return query.Product{}, err
	}

	var variants []query.Variant

	for _, v := range domainEntity.Variants() {
		variants = append(variants, query.Variant{
			ID:    v.ID(),
			Code:  v.Code(),
			Name:  v.Name(),
			Price: v.Price(),
		})
	}

	return query.Product{
		ID:       domainEntity.ID(),
		Name:     domainEntity.Name(),
		Category: domainEntity.CategoryRef(),
		Variants: variants,
	}, nil
}
