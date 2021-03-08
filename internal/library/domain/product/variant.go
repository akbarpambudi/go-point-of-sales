package product

import "go.uber.org/multierr"

type Variant struct {
	id    string
	code  string
	name  string
	price float64
}

func NewVariant(id string, code string, name string, price float64) (variant *Variant, err error) {
	if id == "" {
		err = multierr.Append(err, ErrVariantIDCantBeEmpty)
	}

	if name == "" {
		err = multierr.Append(err, ErrVariantNameCantBeEmpty)
	}

	if price == 0 {
		err = multierr.Append(err, ErrVariantPriceMustGreaterThanZero)
	}

	if err != nil {
		return nil, err
	}

	return &Variant{
		id:    id,
		code:  code,
		name:  name,
		price: price,
	}, nil
}

func (v Variant) ID() string {
	return v.id
}

func (v Variant) Code() string {
	return v.code
}

func (v Variant) Name() string {
	return v.name
}

func (v Variant) Price() float64 {
	return v.price
}
