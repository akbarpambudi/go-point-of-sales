package app

import (
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/command"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/app/query"
)

type Commands struct {
	CreateProduct command.CreateProductHandler
}

type Queries struct {
	GetProductById query.GetProductByIDHandler
}

type Application struct {
	Commands Commands
	Queries  Queries
}
