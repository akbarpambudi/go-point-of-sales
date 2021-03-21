package adapterent

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/query"
)

type ProductsReadModelQueryProjector struct {
	client *ent.Client
}

func NewProductsReadModelQueryProjector(client *ent.Client) *ProductsReadModelQueryProjector {
	return &ProductsReadModelQueryProjector{client: client}
}

func (p ProductsReadModelQueryProjector) GetAllProducts(ctx context.Context) (query.Products, error) {
	result, err := p.client.Product.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var products query.Products
	for _, r := range result {
		p := query.Product{
			ID:       r.ID.String(),
			Name:     r.Name,
			Category: r.CategoryRef,
		}
		variants, err := r.QueryVariants().All(ctx)
		if err != nil {
			return nil, err
		}
		for _, v := range variants {
			p.Variants = append(p.Variants, query.Variant{
				ID:    v.ID.String(),
				Code:  v.Code,
				Name:  v.Name,
				Price: v.Price,
			})
		}

		products = append(products, p)
	}

	return products, nil
}
