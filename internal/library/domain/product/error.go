package product

import "github.com/akbarpambudi/go-point-of-sales/internal/common/errors"

var (
	VariantCreationErrKey              = "error.variant.creation"
	CreationErrKey                     = "error.product.creation"
	ErrProductIDCantBeEmpty            = errors.NewIllegalInputError("err.product.id.shouldNotEmpty", "product id should not empty")
	ErrProductCategoryRefCantBeEmpty   = errors.NewIllegalInputError("err.product.categoryRef.shouldNotEmpty", "product categoryRef should not empty")
	ErrProductNameCantBeEmpty          = errors.NewIllegalInputError("err.product.name.shouldNotEmpty", "product name should not empty")
	ErrProductAtLeastHaveOneVariant    = errors.NewIllegalInputError("err.product.variants.shouldHasOneOrMoreVariant", "product variant should have at least 1 variant")
	ErrVariantIDCantBeEmpty            = errors.NewIllegalInputError("err.variant.id.shouldNotEmpty", "variant id should not empty")
	ErrVariantPriceMustGreaterThanZero = errors.NewIllegalInputError("err.variant.price.shouldMoreThanZero", "variant price should greater than 0")
	ErrVariantNameCantBeEmpty          = errors.NewIllegalInputError("err.variant.name.shouldNotEmpty", "variant name should not empty")
)
