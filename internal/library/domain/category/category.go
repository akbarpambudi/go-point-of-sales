package category

import (
	"go.uber.org/multierr"
)

type Category struct {
	id   string
	name string
}

func NewCategory(id string, name string) (category *Category, err error) {

	if id == "" {
		err = multierr.Append(err, ErrCategoryIDCantBeEmpty)
	}

	if name == "" {
		err = multierr.Append(err, ErrCategoryNameCantBeEmpty)
	}

	if err != nil {
		return nil, err
	}

	return &Category{
		id:   id,
		name: name,
	}, nil
}

func (c Category) ID() string {
	return c.id
}

func (c Category) Name() string {
	return c.name
}
