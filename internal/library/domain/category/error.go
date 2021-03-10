package category

import "github.com/akbarpambudi/go-point-of-sales/internal/common/errors"

var (
	ErrCategoryIDCantBeEmpty   = errors.NewIllegalInputError("err.category.id.shouldNotEmpty", "category id should not empty")
	ErrCategoryNameCantBeEmpty = errors.NewIllegalInputError("err.category.name.shouldNotEmpty", "category name should not empty")
)
