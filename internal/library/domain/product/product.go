package product

import (
	"go.uber.org/multierr"
)

type Product struct {
	id          string
	categoryRef string
	name        string
	variants    []*Variant
}

func NewProduct(id string, categoryRef string, name string, variants []*Variant) (item *Product, err error) {

	if id == "" {
		err = multierr.Append(err, ErrProductIDCantBeEmpty)
	}

	if categoryRef == "" {
		err = multierr.Append(err, ErrProductCategoryRefCantBeEmpty)
	}

	if name == "" {
		err = multierr.Append(err, ErrProductNameCantBeEmpty)
	}

	if variants == nil || len(variants) == 0 {
		err = multierr.Append(err, ErrProductAtLeastHaveOneVariant)
	}

	if err != nil {
		return nil, IllegalProductCreationInputErr(err)
	}

	return &Product{id: id, categoryRef: categoryRef, name: name, variants: variants}, nil
}

func (i Product) ID() string {
	return i.id
}

func (i Product) Name() string {
	return i.name
}

func (i Product) CategoryRef() string {
	return i.categoryRef
}

func (i Product) Variants() []*Variant {
	return i.variants
}

func (i *Product) AddVariant(variant *Variant) {
	i.variants = append(i.variants, variant)
}

func (i *Product) AddManyVariant(variants []*Variant) {
	i.variants = append(i.variants, variants...)
}
