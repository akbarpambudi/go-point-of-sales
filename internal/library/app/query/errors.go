package query

import "github.com/akbarpambudi/go-point-of-sales/internal/common/errors"

var (
	ErrProductResourceNotFound  = errors.NewResourceNotFoundError("err.product.queryById.notFound", "resource not found")
	ErrCategoryResourceNotFound = errors.NewResourceNotFoundError("err.category.queryById.notFound", "resource not found")
)
