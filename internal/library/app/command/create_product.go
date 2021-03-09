package command

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/product"
)

type CreateProduct struct {
	ID          string
	Name        string
	CategoryRef string
	Variants    []*product.Variant
}

type CreateProductHandler struct {
	repo product.Repository
}

func NewCreateProductHandler(repo product.Repository) *CreateProductHandler {
	return &CreateProductHandler{repo: repo}
}

func (c CreateProductHandler) Handle(ctx context.Context, cmd CreateProduct) error {
	ent, err := product.NewProduct(cmd.ID, cmd.CategoryRef, cmd.Name, cmd.Variants)
	if err != nil {
		return err
	}

	return c.repo.Create(ctx, ent)
}
