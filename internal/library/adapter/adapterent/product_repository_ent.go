package adapterent

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"
	entproduct "github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/product"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/product"
	"github.com/google/uuid"
)

type ProductRepository struct {
	client *ent.ProductClient
}

func NewProductRepository(client *ent.ProductClient) *ProductRepository {
	return &ProductRepository{client: client}
}

func (p ProductRepository) Create(ctx context.Context, entity *product.Product) error {

	var variants ent.Variants
	id, err := uuid.Parse(entity.ID())

	if err != nil {
		return err
	}

	for _, v := range entity.Variants() {
		id, err := uuid.Parse(v.ID())
		if err != nil {
			return err
		}
		variants = append(variants, &ent.Variant{
			ID:    id,
			Code:  v.Code(),
			Name:  v.Name(),
			Price: v.Price(),
		})
	}

	_, err = p.client.Create().
		SetID(id).
		SetName(entity.Name()).
		SetCategoryRef(entity.CategoryRef()).
		AddVariants(variants...).
		Save(ctx)
	return err
}

func (p ProductRepository) Update(ctx context.Context, id string, entity *product.Product) error {
	var variants ent.Variants
	_id, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	for _, v := range entity.Variants() {
		id, err := uuid.Parse(v.ID())
		if err != nil {
			return err
		}
		variants = append(variants, &ent.Variant{
			ID:    id,
			Code:  v.Code(),
			Name:  v.Name(),
			Price: v.Price(),
		})
	}
	_, err = p.client.Update().
		SetName(entity.Name()).
		SetCategoryRef(entity.CategoryRef()).
		ClearVariants().
		AddVariants(variants...).
		Where(entproduct.ID(_id)).Save(ctx)

	return err
}

func (p ProductRepository) Load(ctx context.Context, id string) (entity *product.Product, err error) {
	_id, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	result, err := p.client.Query().Where(entproduct.ID(_id)).Where(entproduct.ID(_id)).WithVariants().Only(ctx)

	if err != nil {
		return nil, err
	}

	var variants []*product.Variant

	for _, v := range result.Edges.Variants {
		variant, err := product.NewVariant(v.ID.String(), v.Code, v.Name, v.Price)
		if err != nil {
			return nil, err
		}
		variants = append(variants, variant)
	}

	return product.NewProduct(result.ID.String(), result.CategoryRef, result.Name, variants)
}
