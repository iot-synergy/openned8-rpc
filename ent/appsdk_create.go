// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent/activecodeinfo"
	"github.com/iot-synergy/openned8-rpc/ent/appinfo"
	"github.com/iot-synergy/openned8-rpc/ent/appsdk"
	"github.com/iot-synergy/openned8-rpc/ent/sdkinfo"
)

// AppSdkCreate is the builder for creating a AppSdk entity.
type AppSdkCreate struct {
	config
	mutation *AppSdkMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (asc *AppSdkCreate) SetCreatedAt(t time.Time) *AppSdkCreate {
	asc.mutation.SetCreatedAt(t)
	return asc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (asc *AppSdkCreate) SetNillableCreatedAt(t *time.Time) *AppSdkCreate {
	if t != nil {
		asc.SetCreatedAt(*t)
	}
	return asc
}

// SetUpdatedAt sets the "updated_at" field.
func (asc *AppSdkCreate) SetUpdatedAt(t time.Time) *AppSdkCreate {
	asc.mutation.SetUpdatedAt(t)
	return asc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (asc *AppSdkCreate) SetNillableUpdatedAt(t *time.Time) *AppSdkCreate {
	if t != nil {
		asc.SetUpdatedAt(*t)
	}
	return asc
}

// SetApp sets the "app" field.
func (asc *AppSdkCreate) SetApp(s string) *AppSdkCreate {
	asc.mutation.SetApp(s)
	return asc
}

// SetNillableApp sets the "app" field if the given value is not nil.
func (asc *AppSdkCreate) SetNillableApp(s *string) *AppSdkCreate {
	if s != nil {
		asc.SetApp(*s)
	}
	return asc
}

// SetSdk sets the "sdk" field.
func (asc *AppSdkCreate) SetSdk(s string) *AppSdkCreate {
	asc.mutation.SetSdk(s)
	return asc
}

// SetNillableSdk sets the "sdk" field if the given value is not nil.
func (asc *AppSdkCreate) SetNillableSdk(s *string) *AppSdkCreate {
	if s != nil {
		asc.SetSdk(*s)
	}
	return asc
}

// SetSdkKey sets the "sdk_key" field.
func (asc *AppSdkCreate) SetSdkKey(s string) *AppSdkCreate {
	asc.mutation.SetSdkKey(s)
	return asc
}

// SetID sets the "id" field.
func (asc *AppSdkCreate) SetID(u uuid.UUID) *AppSdkCreate {
	asc.mutation.SetID(u)
	return asc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (asc *AppSdkCreate) SetNillableID(u *uuid.UUID) *AppSdkCreate {
	if u != nil {
		asc.SetID(*u)
	}
	return asc
}

// AddActiveCodeIDs adds the "active_code" edge to the ActiveCodeInfo entity by IDs.
func (asc *AppSdkCreate) AddActiveCodeIDs(ids ...uuid.UUID) *AppSdkCreate {
	asc.mutation.AddActiveCodeIDs(ids...)
	return asc
}

// AddActiveCode adds the "active_code" edges to the ActiveCodeInfo entity.
func (asc *AppSdkCreate) AddActiveCode(a ...*ActiveCodeInfo) *AppSdkCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return asc.AddActiveCodeIDs(ids...)
}

// SetAppInfoID sets the "app_info" edge to the AppInfo entity by ID.
func (asc *AppSdkCreate) SetAppInfoID(id string) *AppSdkCreate {
	asc.mutation.SetAppInfoID(id)
	return asc
}

// SetNillableAppInfoID sets the "app_info" edge to the AppInfo entity by ID if the given value is not nil.
func (asc *AppSdkCreate) SetNillableAppInfoID(id *string) *AppSdkCreate {
	if id != nil {
		asc = asc.SetAppInfoID(*id)
	}
	return asc
}

// SetAppInfo sets the "app_info" edge to the AppInfo entity.
func (asc *AppSdkCreate) SetAppInfo(a *AppInfo) *AppSdkCreate {
	return asc.SetAppInfoID(a.ID)
}

// SetSdkInfoID sets the "sdk_info" edge to the SdkInfo entity by ID.
func (asc *AppSdkCreate) SetSdkInfoID(id string) *AppSdkCreate {
	asc.mutation.SetSdkInfoID(id)
	return asc
}

// SetNillableSdkInfoID sets the "sdk_info" edge to the SdkInfo entity by ID if the given value is not nil.
func (asc *AppSdkCreate) SetNillableSdkInfoID(id *string) *AppSdkCreate {
	if id != nil {
		asc = asc.SetSdkInfoID(*id)
	}
	return asc
}

// SetSdkInfo sets the "sdk_info" edge to the SdkInfo entity.
func (asc *AppSdkCreate) SetSdkInfo(s *SdkInfo) *AppSdkCreate {
	return asc.SetSdkInfoID(s.ID)
}

// Mutation returns the AppSdkMutation object of the builder.
func (asc *AppSdkCreate) Mutation() *AppSdkMutation {
	return asc.mutation
}

// Save creates the AppSdk in the database.
func (asc *AppSdkCreate) Save(ctx context.Context) (*AppSdk, error) {
	asc.defaults()
	return withHooks(ctx, asc.sqlSave, asc.mutation, asc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (asc *AppSdkCreate) SaveX(ctx context.Context) *AppSdk {
	v, err := asc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asc *AppSdkCreate) Exec(ctx context.Context) error {
	_, err := asc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asc *AppSdkCreate) ExecX(ctx context.Context) {
	if err := asc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asc *AppSdkCreate) defaults() {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		v := appsdk.DefaultCreatedAt()
		asc.mutation.SetCreatedAt(v)
	}
	if _, ok := asc.mutation.UpdatedAt(); !ok {
		v := appsdk.DefaultUpdatedAt()
		asc.mutation.SetUpdatedAt(v)
	}
	if _, ok := asc.mutation.ID(); !ok {
		v := appsdk.DefaultID()
		asc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (asc *AppSdkCreate) check() error {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppSdk.created_at"`)}
	}
	if _, ok := asc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppSdk.updated_at"`)}
	}
	if _, ok := asc.mutation.SdkKey(); !ok {
		return &ValidationError{Name: "sdk_key", err: errors.New(`ent: missing required field "AppSdk.sdk_key"`)}
	}
	if v, ok := asc.mutation.SdkKey(); ok {
		if err := appsdk.SdkKeyValidator(v); err != nil {
			return &ValidationError{Name: "sdk_key", err: fmt.Errorf(`ent: validator failed for field "AppSdk.sdk_key": %w`, err)}
		}
	}
	return nil
}

func (asc *AppSdkCreate) sqlSave(ctx context.Context) (*AppSdk, error) {
	if err := asc.check(); err != nil {
		return nil, err
	}
	_node, _spec := asc.createSpec()
	if err := sqlgraph.CreateNode(ctx, asc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	asc.mutation.id = &_node.ID
	asc.mutation.done = true
	return _node, nil
}

func (asc *AppSdkCreate) createSpec() (*AppSdk, *sqlgraph.CreateSpec) {
	var (
		_node = &AppSdk{config: asc.config}
		_spec = sqlgraph.NewCreateSpec(appsdk.Table, sqlgraph.NewFieldSpec(appsdk.FieldID, field.TypeUUID))
	)
	if id, ok := asc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := asc.mutation.CreatedAt(); ok {
		_spec.SetField(appsdk.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := asc.mutation.UpdatedAt(); ok {
		_spec.SetField(appsdk.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := asc.mutation.SdkKey(); ok {
		_spec.SetField(appsdk.FieldSdkKey, field.TypeString, value)
		_node.SdkKey = value
	}
	if nodes := asc.mutation.ActiveCodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   appsdk.ActiveCodeTable,
			Columns: []string{appsdk.ActiveCodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(activecodeinfo.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := asc.mutation.AppInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appsdk.AppInfoTable,
			Columns: []string{appsdk.AppInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appinfo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.App = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := asc.mutation.SdkInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   appsdk.SdkInfoTable,
			Columns: []string{appsdk.SdkInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sdkinfo.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Sdk = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AppSdkCreateBulk is the builder for creating many AppSdk entities in bulk.
type AppSdkCreateBulk struct {
	config
	err      error
	builders []*AppSdkCreate
}

// Save creates the AppSdk entities in the database.
func (ascb *AppSdkCreateBulk) Save(ctx context.Context) ([]*AppSdk, error) {
	if ascb.err != nil {
		return nil, ascb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ascb.builders))
	nodes := make([]*AppSdk, len(ascb.builders))
	mutators := make([]Mutator, len(ascb.builders))
	for i := range ascb.builders {
		func(i int, root context.Context) {
			builder := ascb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppSdkMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ascb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ascb *AppSdkCreateBulk) SaveX(ctx context.Context) []*AppSdk {
	v, err := ascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ascb *AppSdkCreateBulk) Exec(ctx context.Context) error {
	_, err := ascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ascb *AppSdkCreateBulk) ExecX(ctx context.Context) {
	if err := ascb.Exec(ctx); err != nil {
		panic(err)
	}
}
