package category

import "errors"

var (
	ErrCategoryIDCantBeEmpty   = errors.New("category id should not empty")
	ErrCategoryNameCantBeEmpty = errors.New("category name should not empty")
)
