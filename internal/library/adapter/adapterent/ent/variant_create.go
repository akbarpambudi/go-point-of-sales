// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/product"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/variant"
	"github.com/google/uuid"
)

// VariantCreate is the builder for creating a Variant entity.
type VariantCreate struct {
	config
	mutation *VariantMutation
	hooks    []Hook
}

// SetCode sets the "code" field.
func (vc *VariantCreate) SetCode(s string) *VariantCreate {
	vc.mutation.SetCode(s)
	return vc
}

// SetName sets the "name" field.
func (vc *VariantCreate) SetName(s string) *VariantCreate {
	vc.mutation.SetName(s)
	return vc
}

// SetPrice sets the "price" field.
func (vc *VariantCreate) SetPrice(f float64) *VariantCreate {
	vc.mutation.SetPrice(f)
	return vc
}

// SetID sets the "id" field.
func (vc *VariantCreate) SetID(u uuid.UUID) *VariantCreate {
	vc.mutation.SetID(u)
	return vc
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (vc *VariantCreate) SetProductID(id uuid.UUID) *VariantCreate {
	vc.mutation.SetProductID(id)
	return vc
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (vc *VariantCreate) SetNillableProductID(id *uuid.UUID) *VariantCreate {
	if id != nil {
		vc = vc.SetProductID(*id)
	}
	return vc
}

// SetProduct sets the "product" edge to the Product entity.
func (vc *VariantCreate) SetProduct(p *Product) *VariantCreate {
	return vc.SetProductID(p.ID)
}

// Mutation returns the VariantMutation object of the builder.
func (vc *VariantCreate) Mutation() *VariantMutation {
	return vc.mutation
}

// Save creates the Variant in the database.
func (vc *VariantCreate) Save(ctx context.Context) (*Variant, error) {
	var (
		err  error
		node *Variant
	)
	if len(vc.hooks) == 0 {
		if err = vc.check(); err != nil {
			return nil, err
		}
		node, err = vc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VariantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vc.check(); err != nil {
				return nil, err
			}
			vc.mutation = mutation
			node, err = vc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vc.hooks) - 1; i >= 0; i-- {
			mut = vc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VariantCreate) SaveX(ctx context.Context) *Variant {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (vc *VariantCreate) check() error {
	if _, ok := vc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New("ent: missing required field \"code\"")}
	}
	if _, ok := vc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := vc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	return nil
}

func (vc *VariantCreate) sqlSave(ctx context.Context) (*Variant, error) {
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}

func (vc *VariantCreate) createSpec() (*Variant, *sqlgraph.CreateSpec) {
	var (
		_node = &Variant{config: vc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: variant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: variant.FieldID,
			},
		}
	)
	if id, ok := vc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: variant.FieldCode,
		})
		_node.Code = value
	}
	if value, ok := vc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: variant.FieldName,
		})
		_node.Name = value
	}
	if value, ok := vc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: variant.FieldPrice,
		})
		_node.Price = value
	}
	if nodes := vc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   variant.ProductTable,
			Columns: []string{variant.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VariantCreateBulk is the builder for creating many Variant entities in bulk.
type VariantCreateBulk struct {
	config
	builders []*VariantCreate
}

// Save creates the Variant entities in the database.
func (vcb *VariantCreateBulk) Save(ctx context.Context) ([]*Variant, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Variant, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VariantMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VariantCreateBulk) SaveX(ctx context.Context) []*Variant {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}