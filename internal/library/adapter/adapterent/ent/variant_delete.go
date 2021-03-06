// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/predicate"
	"github.com/akbarpambudi/go-point-of-sales/internal/library/adapter/adapterent/ent/variant"
)

// VariantDelete is the builder for deleting a Variant entity.
type VariantDelete struct {
	config
	hooks    []Hook
	mutation *VariantMutation
}

// Where adds a new predicate to the VariantDelete builder.
func (vd *VariantDelete) Where(ps ...predicate.Variant) *VariantDelete {
	vd.mutation.predicates = append(vd.mutation.predicates, ps...)
	return vd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vd *VariantDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vd.hooks) == 0 {
		affected, err = vd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VariantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vd.mutation = mutation
			affected, err = vd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vd.hooks) - 1; i >= 0; i-- {
			mut = vd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vd *VariantDelete) ExecX(ctx context.Context) int {
	n, err := vd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vd *VariantDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: variant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: variant.FieldID,
			},
		},
	}
	if ps := vd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vd.driver, _spec)
}

// VariantDeleteOne is the builder for deleting a single Variant entity.
type VariantDeleteOne struct {
	vd *VariantDelete
}

// Exec executes the deletion query.
func (vdo *VariantDeleteOne) Exec(ctx context.Context) error {
	n, err := vdo.vd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{variant.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vdo *VariantDeleteOne) ExecX(ctx context.Context) {
	vdo.vd.ExecX(ctx)
}
