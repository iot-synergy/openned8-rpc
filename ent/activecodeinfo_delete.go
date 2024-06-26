// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/openned8-rpc/ent/activecodeinfo"
	"github.com/iot-synergy/openned8-rpc/ent/predicate"
)

// ActiveCodeInfoDelete is the builder for deleting a ActiveCodeInfo entity.
type ActiveCodeInfoDelete struct {
	config
	hooks    []Hook
	mutation *ActiveCodeInfoMutation
}

// Where appends a list predicates to the ActiveCodeInfoDelete builder.
func (acid *ActiveCodeInfoDelete) Where(ps ...predicate.ActiveCodeInfo) *ActiveCodeInfoDelete {
	acid.mutation.Where(ps...)
	return acid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (acid *ActiveCodeInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, acid.sqlExec, acid.mutation, acid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (acid *ActiveCodeInfoDelete) ExecX(ctx context.Context) int {
	n, err := acid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (acid *ActiveCodeInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(activecodeinfo.Table, sqlgraph.NewFieldSpec(activecodeinfo.FieldID, field.TypeUUID))
	if ps := acid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, acid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	acid.mutation.done = true
	return affected, err
}

// ActiveCodeInfoDeleteOne is the builder for deleting a single ActiveCodeInfo entity.
type ActiveCodeInfoDeleteOne struct {
	acid *ActiveCodeInfoDelete
}

// Where appends a list predicates to the ActiveCodeInfoDelete builder.
func (acido *ActiveCodeInfoDeleteOne) Where(ps ...predicate.ActiveCodeInfo) *ActiveCodeInfoDeleteOne {
	acido.acid.mutation.Where(ps...)
	return acido
}

// Exec executes the deletion query.
func (acido *ActiveCodeInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := acido.acid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{activecodeinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (acido *ActiveCodeInfoDeleteOne) ExecX(ctx context.Context) {
	if err := acido.Exec(ctx); err != nil {
		panic(err)
	}
}
