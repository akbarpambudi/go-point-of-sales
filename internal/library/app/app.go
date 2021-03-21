package app

import (
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/command"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/query"
)

type Commands struct {
	CreateProduct  command.CreateProductHandler
	CreateCategory command.CreateCategoryHandler
}

type Queries struct {
	GetProductById  query.GetProductByIDHandler
	GetCategoryById query.GetCategoryByIDHandler
	GetAllProducts  query.GetAllProductsHandler
}

type Application struct {
	Commands Commands
	Queries  Queries
}
