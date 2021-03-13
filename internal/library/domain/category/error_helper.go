package category

import "github.com/akbarpambudi/go-point-of-sales/internal/common/errors"

func IllegalCreationInputError(err error) error {
	return errors.WrapMultiErr(errors.ErrorTypeIllegalInputError, CreationErrKey, "invalid input data", err)
}
