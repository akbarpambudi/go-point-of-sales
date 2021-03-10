package command

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/domain/category"
)

type CreateCategory struct {
	ID   string
	Name string
}

type CreateCategoryHandler interface {
	Handle(ctx context.Context, cmd CreateCategory) error
}

type CreateCategoryHandlerImpl struct {
	repository category.Repository
}

func NewCreateCategoryHandlerImpl(repository category.Repository) *CreateCategoryHandlerImpl {
	return &CreateCategoryHandlerImpl{repository: repository}
}

func (c CreateCategoryHandlerImpl) Handle(ctx context.Context, cmd CreateCategory) error {

	entity, err := category.NewCategory(cmd.ID, cmd.Name)
	if err != nil {
		return err
	}

	return c.repository.Create(ctx, entity)
}
