package product

import "github.com/pkg/errors"

var (
	ErrProductIDCantBeEmpty            = errors.New("product id should not empty")
	ErrProductCategoryRefCantBeEmpty   = errors.New("product categoryRef should not empty")
	ErrProductNameCantBeEmpty          = errors.New("product name should not empty")
	ErrProductAtLeastHaveOneVariant    = errors.New("product variant should have at least 1 variant")
	ErrVariantIDCantBeEmpty            = errors.New("variant id should not empty")
	ErrVariantPriceMustGreaterThanZero = errors.New("variant price should greater than 0")
	ErrVariantNameCantBeEmpty          = errors.New("variant name should not empty")
	ErrCategoryIDCantBeEmpty           = errors.New("category id should not empty")
	ErrCategoryNameCantBeEmpty         = errors.New("category name should not empty")
)
