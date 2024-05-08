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
	"github.com/iot-synergy/openned8-rpc/ent/appsdk"
	"github.com/iot-synergy/openned8-rpc/ent/sdkinfo"
)

// SdkInfoCreate is the builder for creating a SdkInfo entity.
type SdkInfoCreate struct {
	config
	mutation *SdkInfoMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sic *SdkInfoCreate) SetCreatedAt(t time.Time) *SdkInfoCreate {
	sic.mutation.SetCreatedAt(t)
	return sic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sic *SdkInfoCreate) SetNillableCreatedAt(t *time.Time) *SdkInfoCreate {
	if t != nil {
		sic.SetCreatedAt(*t)
	}
	return sic
}

// SetUpdatedAt sets the "updated_at" field.
func (sic *SdkInfoCreate) SetUpdatedAt(t time.Time) *SdkInfoCreate {
	sic.mutation.SetUpdatedAt(t)
	return sic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sic *SdkInfoCreate) SetNillableUpdatedAt(t *time.Time) *SdkInfoCreate {
	if t != nil {
		sic.SetUpdatedAt(*t)
	}
	return sic
}

// SetStatus sets the "status" field.
func (sic *SdkInfoCreate) SetStatus(u uint8) *SdkInfoCreate {
	sic.mutation.SetStatus(u)
	return sic
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sic *SdkInfoCreate) SetNillableStatus(u *uint8) *SdkInfoCreate {
	if u != nil {
		sic.SetStatus(*u)
	}
	return sic
}

// SetName sets the "name" field.
func (sic *SdkInfoCreate) SetName(s string) *SdkInfoCreate {
	sic.mutation.SetName(s)
	return sic
}

// SetAvatar sets the "avatar" field.
func (sic *SdkInfoCreate) SetAvatar(s string) *SdkInfoCreate {
	sic.mutation.SetAvatar(s)
	return sic
}

// SetDesc sets the "desc" field.
func (sic *SdkInfoCreate) SetDesc(i int64) *SdkInfoCreate {
	sic.mutation.SetDesc(i)
	return sic
}

// SetDownloadURL sets the "download_url" field.
func (sic *SdkInfoCreate) SetDownloadURL(s string) *SdkInfoCreate {
	sic.mutation.SetDownloadURL(s)
	return sic
}

// SetID sets the "id" field.
func (sic *SdkInfoCreate) SetID(u uuid.UUID) *SdkInfoCreate {
	sic.mutation.SetID(u)
	return sic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sic *SdkInfoCreate) SetNillableID(u *uuid.UUID) *SdkInfoCreate {
	if u != nil {
		sic.SetID(*u)
	}
	return sic
}

// AddAppSdkIDs adds the "app_sdk" edge to the AppSdk entity by IDs.
func (sic *SdkInfoCreate) AddAppSdkIDs(ids ...uuid.UUID) *SdkInfoCreate {
	sic.mutation.AddAppSdkIDs(ids...)
	return sic
}

// AddAppSdk adds the "app_sdk" edges to the AppSdk entity.
func (sic *SdkInfoCreate) AddAppSdk(a ...*AppSdk) *SdkInfoCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sic.AddAppSdkIDs(ids...)
}

// Mutation returns the SdkInfoMutation object of the builder.
func (sic *SdkInfoCreate) Mutation() *SdkInfoMutation {
	return sic.mutation
}

// Save creates the SdkInfo in the database.
func (sic *SdkInfoCreate) Save(ctx context.Context) (*SdkInfo, error) {
	sic.defaults()
	return withHooks(ctx, sic.sqlSave, sic.mutation, sic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sic *SdkInfoCreate) SaveX(ctx context.Context) *SdkInfo {
	v, err := sic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sic *SdkInfoCreate) Exec(ctx context.Context) error {
	_, err := sic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sic *SdkInfoCreate) ExecX(ctx context.Context) {
	if err := sic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sic *SdkInfoCreate) defaults() {
	if _, ok := sic.mutation.CreatedAt(); !ok {
		v := sdkinfo.DefaultCreatedAt()
		sic.mutation.SetCreatedAt(v)
	}
	if _, ok := sic.mutation.UpdatedAt(); !ok {
		v := sdkinfo.DefaultUpdatedAt()
		sic.mutation.SetUpdatedAt(v)
	}
	if _, ok := sic.mutation.Status(); !ok {
		v := sdkinfo.DefaultStatus
		sic.mutation.SetStatus(v)
	}
	if _, ok := sic.mutation.ID(); !ok {
		v := sdkinfo.DefaultID()
		sic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sic *SdkInfoCreate) check() error {
	if _, ok := sic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SdkInfo.created_at"`)}
	}
	if _, ok := sic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SdkInfo.updated_at"`)}
	}
	if _, ok := sic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "SdkInfo.name"`)}
	}
	if _, ok := sic.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "SdkInfo.avatar"`)}
	}
	if _, ok := sic.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "SdkInfo.desc"`)}
	}
	if _, ok := sic.mutation.DownloadURL(); !ok {
		return &ValidationError{Name: "download_url", err: errors.New(`ent: missing required field "SdkInfo.download_url"`)}
	}
	return nil
}

func (sic *SdkInfoCreate) sqlSave(ctx context.Context) (*SdkInfo, error) {
	if err := sic.check(); err != nil {
		return nil, err
	}
	_node, _spec := sic.createSpec()
	if err := sqlgraph.CreateNode(ctx, sic.driver, _spec); err != nil {
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
	sic.mutation.id = &_node.ID
	sic.mutation.done = true
	return _node, nil
}

func (sic *SdkInfoCreate) createSpec() (*SdkInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &SdkInfo{config: sic.config}
		_spec = sqlgraph.NewCreateSpec(sdkinfo.Table, sqlgraph.NewFieldSpec(sdkinfo.FieldID, field.TypeUUID))
	)
	if id, ok := sic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sic.mutation.CreatedAt(); ok {
		_spec.SetField(sdkinfo.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sic.mutation.UpdatedAt(); ok {
		_spec.SetField(sdkinfo.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sic.mutation.Status(); ok {
		_spec.SetField(sdkinfo.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := sic.mutation.Name(); ok {
		_spec.SetField(sdkinfo.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sic.mutation.Avatar(); ok {
		_spec.SetField(sdkinfo.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := sic.mutation.Desc(); ok {
		_spec.SetField(sdkinfo.FieldDesc, field.TypeInt64, value)
		_node.Desc = value
	}
	if value, ok := sic.mutation.DownloadURL(); ok {
		_spec.SetField(sdkinfo.FieldDownloadURL, field.TypeString, value)
		_node.DownloadURL = value
	}
	if nodes := sic.mutation.AppSdkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   sdkinfo.AppSdkTable,
			Columns: []string{sdkinfo.AppSdkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appsdk.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SdkInfoCreateBulk is the builder for creating many SdkInfo entities in bulk.
type SdkInfoCreateBulk struct {
	config
	err      error
	builders []*SdkInfoCreate
}

// Save creates the SdkInfo entities in the database.
func (sicb *SdkInfoCreateBulk) Save(ctx context.Context) ([]*SdkInfo, error) {
	if sicb.err != nil {
		return nil, sicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(sicb.builders))
	nodes := make([]*SdkInfo, len(sicb.builders))
	mutators := make([]Mutator, len(sicb.builders))
	for i := range sicb.builders {
		func(i int, root context.Context) {
			builder := sicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SdkInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, sicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, sicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, sicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (sicb *SdkInfoCreateBulk) SaveX(ctx context.Context) []*SdkInfo {
	v, err := sicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sicb *SdkInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := sicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sicb *SdkInfoCreateBulk) ExecX(ctx context.Context) {
	if err := sicb.Exec(ctx); err != nil {
		panic(err)
	}
}