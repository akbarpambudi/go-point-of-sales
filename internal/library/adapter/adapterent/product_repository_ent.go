package adapterent

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"
	entproduct "github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/product"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/product"
	"github.com/google/uuid"
)

type ProductRepository struct {
	client *ent.Client
}

func NewProductRepository(client *ent.Client) *ProductRepository {
	return &ProductRepository{client: client}
}

func (p ProductRepository) Create(ctx context.Context, entity *product.Product) (err error) {
	tx, err := p.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if tx != nil {
				err = tx.Rollback()
			}
		}
	}()

	var variantCreationQueries []*ent.VariantCreate
	id, err := uuid.Parse(entity.ID())

	if err != nil {
		return err
	}

	for _, v := range entity.Variants() {
		id, err := uuid.Parse(v.ID())
		if err != nil {
			return err
		}
		query := tx.Variant.Create().SetID(id).SetName(v.Name()).SetCode(v.Code()).SetPrice(v.Price())
		variantCreationQueries = append(variantCreationQueries, query)
	}

	variants, err := tx.Variant.CreateBulk(variantCreationQueries...).Save(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Product.Create().
		SetID(id).
		SetName(entity.Name()).
		SetCategoryRef(entity.CategoryRef()).
		AddVariants(variants...).
		Save(ctx)

	if err != nil {
		return err
	}

	return tx.Commit()
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
	_, err = p.client.Product.Update().
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
	result, err := p.client.Product.Query().Where(entproduct.ID(_id)).Where(entproduct.ID(_id)).WithVariants().Only(ctx)

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
