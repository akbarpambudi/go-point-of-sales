package command

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/product"
)

type CreateProduct struct {
	ID          string
	Name        string
	CategoryRef string
	Variants    []ProductVariantDTO
}

type CreateProductHandler interface {
	Handle(ctx context.Context, cmd CreateProduct) error
}

type CreateProductHandlerImpl struct {
	repo product.Repository
}

func NewCreateProductHandlerImpl(repo product.Repository) *CreateProductHandlerImpl {
	return &CreateProductHandlerImpl{repo: repo}
}

func (c CreateProductHandlerImpl) Handle(ctx context.Context, cmd CreateProduct) error {

	var productVariants []*product.Variant

	for _, v := range cmd.Variants {
		productVariant, err := product.NewVariant(v.ID, v.Code, v.Name, v.Price)
		if err != nil {
			return err
		}

		productVariants = append(productVariants, productVariant)
	}

	ent, err := product.NewProduct(cmd.ID, cmd.CategoryRef, cmd.Name, productVariants)
	if err != nil {
		return err
	}

	return c.repo.Create(ctx, ent)
}
