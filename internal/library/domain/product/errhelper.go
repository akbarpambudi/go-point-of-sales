package product

import "github.com/akbarpambudi/go-point-of-sales/internal/common/errors"

func IllegalProductCreationInputErr(err error) error {
	return errors.WrapMultiErr(errors.ErrorTypeIllegalInputError, CreationErrKey, "invalid input data", err)
}

func IllegalVariantCreationInputErr(err error) error {
	return errors.WrapMultiErr(errors.ErrorTypeIllegalInputError, VariantCreationErrKey, "invalid input data", err)
}
