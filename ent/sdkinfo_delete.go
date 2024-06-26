// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/openned8-rpc/ent/predicate"
	"github.com/iot-synergy/openned8-rpc/ent/sdkinfo"
)

// SdkInfoDelete is the builder for deleting a SdkInfo entity.
type SdkInfoDelete struct {
	config
	hooks    []Hook
	mutation *SdkInfoMutation
}

// Where appends a list predicates to the SdkInfoDelete builder.
func (sid *SdkInfoDelete) Where(ps ...predicate.SdkInfo) *SdkInfoDelete {
	sid.mutation.Where(ps...)
	return sid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sid *SdkInfoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sid.sqlExec, sid.mutation, sid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sid *SdkInfoDelete) ExecX(ctx context.Context) int {
	n, err := sid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sid *SdkInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sdkinfo.Table, sqlgraph.NewFieldSpec(sdkinfo.FieldID, field.TypeString))
	if ps := sid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sid.mutation.done = true
	return affected, err
}

// SdkInfoDeleteOne is the builder for deleting a single SdkInfo entity.
type SdkInfoDeleteOne struct {
	sid *SdkInfoDelete
}

// Where appends a list predicates to the SdkInfoDelete builder.
func (sido *SdkInfoDeleteOne) Where(ps ...predicate.SdkInfo) *SdkInfoDeleteOne {
	sido.sid.mutation.Where(ps...)
	return sido
}

// Exec executes the deletion query.
func (sido *SdkInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := sido.sid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sdkinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sido *SdkInfoDeleteOne) ExecX(ctx context.Context) {
	if err := sido.Exec(ctx); err != nil {
		panic(err)
	}
}
