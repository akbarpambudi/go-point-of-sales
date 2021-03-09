package app

import "github.com/akbarpambudi/go-point-of-sales/internal/library/app/command"

type Commands struct {
	CreateProduct command.CreateProductHandler
}

type Application struct {
	Commands Commands
}
