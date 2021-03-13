package service

import (
	"context"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/command"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/query"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type ApplicationOptions struct {
	Client *ent.Client
}

func NewApplication(ctx context.Context, options ApplicationOptions) (app.Application, func(), error) {
	noopCleansingFunc := func() {}
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	if options.Client != nil {
		client = options.Client
	}

	if err != nil {
		return app.Application{}, noopCleansingFunc, err
	}

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	productRepository := adapterent.NewProductRepository(client)
	categoryRepository := adapterent.NewCategoryRepositoryEnt(client)

	return app.Application{
			Commands: app.Commands{
				CreateProduct:  command.NewCreateProductHandlerImpl(productRepository),
				CreateCategory: command.NewCreateCategoryHandlerImpl(categoryRepository),
			},
			Queries: app.Queries{
				GetProductById: query.NewGetProductByIDHandlerImpl(productRepository),
			},
		}, func() {
			err := client.Close()
			if err != nil {
				log.Fatalf("unexpected error happened while closing connection: %v", err)
			}
		}, nil
}
