// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/predicate"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/product"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/variant"
	"github.com/google/uuid"
)

// VariantUpdate is the builder for updating Variant entities.
type VariantUpdate struct {
	config
	hooks    []Hook
	mutation *VariantMutation
}

// Where adds a new predicate for the VariantUpdate builder.
func (vu *VariantUpdate) Where(ps ...predicate.Variant) *VariantUpdate {
	vu.mutation.predicates = append(vu.mutation.predicates, ps...)
	return vu
}

// SetCode sets the "code" field.
func (vu *VariantUpdate) SetCode(s string) *VariantUpdate {
	vu.mutation.SetCode(s)
	return vu
}

// SetName sets the "name" field.
func (vu *VariantUpdate) SetName(s string) *VariantUpdate {
	vu.mutation.SetName(s)
	return vu
}

// SetPrice sets the "price" field.
func (vu *VariantUpdate) SetPrice(f float64) *VariantUpdate {
	vu.mutation.ResetPrice()
	vu.mutation.SetPrice(f)
	return vu
}

// AddPrice adds f to the "price" field.
func (vu *VariantUpdate) AddPrice(f float64) *VariantUpdate {
	vu.mutation.AddPrice(f)
	return vu
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (vu *VariantUpdate) SetProductID(id uuid.UUID) *VariantUpdate {
	vu.mutation.SetProductID(id)
	return vu
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (vu *VariantUpdate) SetNillableProductID(id *uuid.UUID) *VariantUpdate {
	if id != nil {
		vu = vu.SetProductID(*id)
	}
	return vu
}

// SetProduct sets the "product" edge to the Product entity.
func (vu *VariantUpdate) SetProduct(p *Product) *VariantUpdate {
	return vu.SetProductID(p.ID)
}

// Mutation returns the VariantMutation object of the builder.
func (vu *VariantUpdate) Mutation() *VariantMutation {
	return vu.mutation
}

// ClearProduct clears the "product" edge to the Product entity.
func (vu *VariantUpdate) ClearProduct() *VariantUpdate {
	vu.mutation.ClearProduct()
	return vu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VariantUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vu.hooks) == 0 {
		affected, err = vu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VariantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vu.mutation = mutation
			affected, err = vu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vu.hooks) - 1; i >= 0; i-- {
			mut = vu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VariantUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VariantUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VariantUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vu *VariantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   variant.Table,
			Columns: variant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: variant.FieldID,
			},
		},
	}
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: variant.FieldCode,
		})
	}
	if value, ok := vu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: variant.FieldName,
		})
	}
	if value, ok := vu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: variant.FieldPrice,
		})
	}
	if value, ok := vu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: variant.FieldPrice,
		})
	}
	if vu.mutation.ProductCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vu.mutation.ProductIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{variant.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// VariantUpdateOne is the builder for updating a single Variant entity.
type VariantUpdateOne struct {
	config
	hooks    []Hook
	mutation *VariantMutation
}

// SetCode sets the "code" field.
func (vuo *VariantUpdateOne) SetCode(s string) *VariantUpdateOne {
	vuo.mutation.SetCode(s)
	return vuo
}

// SetName sets the "name" field.
func (vuo *VariantUpdateOne) SetName(s string) *VariantUpdateOne {
	vuo.mutation.SetName(s)
	return vuo
}

// SetPrice sets the "price" field.
func (vuo *VariantUpdateOne) SetPrice(f float64) *VariantUpdateOne {
	vuo.mutation.ResetPrice()
	vuo.mutation.SetPrice(f)
	return vuo
}

// AddPrice adds f to the "price" field.
func (vuo *VariantUpdateOne) AddPrice(f float64) *VariantUpdateOne {
	vuo.mutation.AddPrice(f)
	return vuo
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (vuo *VariantUpdateOne) SetProductID(id uuid.UUID) *VariantUpdateOne {
	vuo.mutation.SetProductID(id)
	return vuo
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (vuo *VariantUpdateOne) SetNillableProductID(id *uuid.UUID) *VariantUpdateOne {
	if id != nil {
		vuo = vuo.SetProductID(*id)
	}
	return vuo
}

// SetProduct sets the "product" edge to the Product entity.
func (vuo *VariantUpdateOne) SetProduct(p *Product) *VariantUpdateOne {
	return vuo.SetProductID(p.ID)
}

// Mutation returns the VariantMutation object of the builder.
func (vuo *VariantUpdateOne) Mutation() *VariantMutation {
	return vuo.mutation
}

// ClearProduct clears the "product" edge to the Product entity.
func (vuo *VariantUpdateOne) ClearProduct() *VariantUpdateOne {
	vuo.mutation.ClearProduct()
	return vuo
}

// Save executes the query and returns the updated Variant entity.
func (vuo *VariantUpdateOne) Save(ctx context.Context) (*Variant, error) {
	var (
		err  error
		node *Variant
	)
	if len(vuo.hooks) == 0 {
		node, err = vuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VariantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vuo.mutation = mutation
			node, err = vuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vuo.hooks) - 1; i >= 0; i-- {
			mut = vuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VariantUpdateOne) SaveX(ctx context.Context) *Variant {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VariantUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VariantUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vuo *VariantUpdateOne) sqlSave(ctx context.Context) (_node *Variant, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   variant.Table,
			Columns: variant.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: variant.FieldID,
			},
		},
	}
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Variant.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: variant.FieldCode,
		})
	}
	if value, ok := vuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: variant.FieldName,
		})
	}
	if value, ok := vuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: variant.FieldPrice,
		})
	}
	if value, ok := vuo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: variant.FieldPrice,
		})
	}
	if vuo.mutation.ProductCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuo.mutation.ProductIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Variant{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{variant.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
