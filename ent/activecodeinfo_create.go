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
	"github.com/iot-synergy/openned8-rpc/ent/appsdk"
)

// ActiveCodeInfoCreate is the builder for creating a ActiveCodeInfo entity.
type ActiveCodeInfoCreate struct {
	config
	mutation *ActiveCodeInfoMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (acic *ActiveCodeInfoCreate) SetCreatedAt(t time.Time) *ActiveCodeInfoCreate {
	acic.mutation.SetCreatedAt(t)
	return acic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (acic *ActiveCodeInfoCreate) SetNillableCreatedAt(t *time.Time) *ActiveCodeInfoCreate {
	if t != nil {
		acic.SetCreatedAt(*t)
	}
	return acic
}

// SetUpdatedAt sets the "updated_at" field.
func (acic *ActiveCodeInfoCreate) SetUpdatedAt(t time.Time) *ActiveCodeInfoCreate {
	acic.mutation.SetUpdatedAt(t)
	return acic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (acic *ActiveCodeInfoCreate) SetNillableUpdatedAt(t *time.Time) *ActiveCodeInfoCreate {
	if t != nil {
		acic.SetUpdatedAt(*t)
	}
	return acic
}

// SetStatus sets the "status" field.
func (acic *ActiveCodeInfoCreate) SetStatus(u uint8) *ActiveCodeInfoCreate {
	acic.mutation.SetStatus(u)
	return acic
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (acic *ActiveCodeInfoCreate) SetNillableStatus(u *uint8) *ActiveCodeInfoCreate {
	if u != nil {
		acic.SetStatus(*u)
	}
	return acic
}

// SetActiveKey sets the "active_key" field.
func (acic *ActiveCodeInfoCreate) SetActiveKey(s string) *ActiveCodeInfoCreate {
	acic.mutation.SetActiveKey(s)
	return acic
}

// SetUserID sets the "user_id" field.
func (acic *ActiveCodeInfoCreate) SetUserID(s string) *ActiveCodeInfoCreate {
	acic.mutation.SetUserID(s)
	return acic
}

// SetAppID sets the "app_id" field.
func (acic *ActiveCodeInfoCreate) SetAppID(s string) *ActiveCodeInfoCreate {
	acic.mutation.SetAppID(s)
	return acic
}

// SetActiveIP sets the "active_ip" field.
func (acic *ActiveCodeInfoCreate) SetActiveIP(s string) *ActiveCodeInfoCreate {
	acic.mutation.SetActiveIP(s)
	return acic
}

// SetDeviceSn sets the "device_sn" field.
func (acic *ActiveCodeInfoCreate) SetDeviceSn(s string) *ActiveCodeInfoCreate {
	acic.mutation.SetDeviceSn(s)
	return acic
}

// SetDeviceMAC sets the "device_mac" field.
func (acic *ActiveCodeInfoCreate) SetDeviceMAC(s string) *ActiveCodeInfoCreate {
	acic.mutation.SetDeviceMAC(s)
	return acic
}

// SetDeviceIdentity sets the "device_identity" field.
func (acic *ActiveCodeInfoCreate) SetDeviceIdentity(s string) *ActiveCodeInfoCreate {
	acic.mutation.SetDeviceIdentity(s)
	return acic
}

// SetActiveDate sets the "active_date" field.
func (acic *ActiveCodeInfoCreate) SetActiveDate(t time.Time) *ActiveCodeInfoCreate {
	acic.mutation.SetActiveDate(t)
	return acic
}

// SetNillableActiveDate sets the "active_date" field if the given value is not nil.
func (acic *ActiveCodeInfoCreate) SetNillableActiveDate(t *time.Time) *ActiveCodeInfoCreate {
	if t != nil {
		acic.SetActiveDate(*t)
	}
	return acic
}

// SetActiveType sets the "active_type" field.
func (acic *ActiveCodeInfoCreate) SetActiveType(i int64) *ActiveCodeInfoCreate {
	acic.mutation.SetActiveType(i)
	return acic
}

// SetNillableActiveType sets the "active_type" field if the given value is not nil.
func (acic *ActiveCodeInfoCreate) SetNillableActiveType(i *int64) *ActiveCodeInfoCreate {
	if i != nil {
		acic.SetActiveType(*i)
	}
	return acic
}

// SetActiveFile sets the "active_file" field.
func (acic *ActiveCodeInfoCreate) SetActiveFile(s string) *ActiveCodeInfoCreate {
	acic.mutation.SetActiveFile(s)
	return acic
}

// SetVersion sets the "version" field.
func (acic *ActiveCodeInfoCreate) SetVersion(s string) *ActiveCodeInfoCreate {
	acic.mutation.SetVersion(s)
	return acic
}

// SetStartDate sets the "start_date" field.
func (acic *ActiveCodeInfoCreate) SetStartDate(t time.Time) *ActiveCodeInfoCreate {
	acic.mutation.SetStartDate(t)
	return acic
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (acic *ActiveCodeInfoCreate) SetNillableStartDate(t *time.Time) *ActiveCodeInfoCreate {
	if t != nil {
		acic.SetStartDate(*t)
	}
	return acic
}

// SetExpireDate sets the "expire_date" field.
func (acic *ActiveCodeInfoCreate) SetExpireDate(t time.Time) *ActiveCodeInfoCreate {
	acic.mutation.SetExpireDate(t)
	return acic
}

// SetAppSdkID sets the "app_sdk_id" field.
func (acic *ActiveCodeInfoCreate) SetAppSdkID(u uuid.UUID) *ActiveCodeInfoCreate {
	acic.mutation.SetAppSdkID(u)
	return acic
}

// SetNillableAppSdkID sets the "app_sdk_id" field if the given value is not nil.
func (acic *ActiveCodeInfoCreate) SetNillableAppSdkID(u *uuid.UUID) *ActiveCodeInfoCreate {
	if u != nil {
		acic.SetAppSdkID(*u)
	}
	return acic
}

// SetID sets the "id" field.
func (acic *ActiveCodeInfoCreate) SetID(u uuid.UUID) *ActiveCodeInfoCreate {
	acic.mutation.SetID(u)
	return acic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (acic *ActiveCodeInfoCreate) SetNillableID(u *uuid.UUID) *ActiveCodeInfoCreate {
	if u != nil {
		acic.SetID(*u)
	}
	return acic
}

// SetAppSdk sets the "app_sdk" edge to the AppSdk entity.
func (acic *ActiveCodeInfoCreate) SetAppSdk(a *AppSdk) *ActiveCodeInfoCreate {
	return acic.SetAppSdkID(a.ID)
}

// Mutation returns the ActiveCodeInfoMutation object of the builder.
func (acic *ActiveCodeInfoCreate) Mutation() *ActiveCodeInfoMutation {
	return acic.mutation
}

// Save creates the ActiveCodeInfo in the database.
func (acic *ActiveCodeInfoCreate) Save(ctx context.Context) (*ActiveCodeInfo, error) {
	acic.defaults()
	return withHooks(ctx, acic.sqlSave, acic.mutation, acic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (acic *ActiveCodeInfoCreate) SaveX(ctx context.Context) *ActiveCodeInfo {
	v, err := acic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acic *ActiveCodeInfoCreate) Exec(ctx context.Context) error {
	_, err := acic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acic *ActiveCodeInfoCreate) ExecX(ctx context.Context) {
	if err := acic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acic *ActiveCodeInfoCreate) defaults() {
	if _, ok := acic.mutation.CreatedAt(); !ok {
		v := activecodeinfo.DefaultCreatedAt()
		acic.mutation.SetCreatedAt(v)
	}
	if _, ok := acic.mutation.UpdatedAt(); !ok {
		v := activecodeinfo.DefaultUpdatedAt()
		acic.mutation.SetUpdatedAt(v)
	}
	if _, ok := acic.mutation.Status(); !ok {
		v := activecodeinfo.DefaultStatus
		acic.mutation.SetStatus(v)
	}
	if _, ok := acic.mutation.ActiveDate(); !ok {
		v := activecodeinfo.DefaultActiveDate
		acic.mutation.SetActiveDate(v)
	}
	if _, ok := acic.mutation.ActiveType(); !ok {
		v := activecodeinfo.DefaultActiveType
		acic.mutation.SetActiveType(v)
	}
	if _, ok := acic.mutation.StartDate(); !ok {
		v := activecodeinfo.DefaultStartDate
		acic.mutation.SetStartDate(v)
	}
	if _, ok := acic.mutation.ID(); !ok {
		v := activecodeinfo.DefaultID()
		acic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acic *ActiveCodeInfoCreate) check() error {
	if _, ok := acic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ActiveCodeInfo.created_at"`)}
	}
	if _, ok := acic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ActiveCodeInfo.updated_at"`)}
	}
	if _, ok := acic.mutation.ActiveKey(); !ok {
		return &ValidationError{Name: "active_key", err: errors.New(`ent: missing required field "ActiveCodeInfo.active_key"`)}
	}
	if _, ok := acic.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "ActiveCodeInfo.user_id"`)}
	}
	if _, ok := acic.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "ActiveCodeInfo.app_id"`)}
	}
	if _, ok := acic.mutation.ActiveIP(); !ok {
		return &ValidationError{Name: "active_ip", err: errors.New(`ent: missing required field "ActiveCodeInfo.active_ip"`)}
	}
	if _, ok := acic.mutation.DeviceSn(); !ok {
		return &ValidationError{Name: "device_sn", err: errors.New(`ent: missing required field "ActiveCodeInfo.device_sn"`)}
	}
	if _, ok := acic.mutation.DeviceMAC(); !ok {
		return &ValidationError{Name: "device_mac", err: errors.New(`ent: missing required field "ActiveCodeInfo.device_mac"`)}
	}
	if _, ok := acic.mutation.DeviceIdentity(); !ok {
		return &ValidationError{Name: "device_identity", err: errors.New(`ent: missing required field "ActiveCodeInfo.device_identity"`)}
	}
	if _, ok := acic.mutation.ActiveDate(); !ok {
		return &ValidationError{Name: "active_date", err: errors.New(`ent: missing required field "ActiveCodeInfo.active_date"`)}
	}
	if _, ok := acic.mutation.ActiveType(); !ok {
		return &ValidationError{Name: "active_type", err: errors.New(`ent: missing required field "ActiveCodeInfo.active_type"`)}
	}
	if _, ok := acic.mutation.ActiveFile(); !ok {
		return &ValidationError{Name: "active_file", err: errors.New(`ent: missing required field "ActiveCodeInfo.active_file"`)}
	}
	if _, ok := acic.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "ActiveCodeInfo.version"`)}
	}
	if _, ok := acic.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New(`ent: missing required field "ActiveCodeInfo.start_date"`)}
	}
	if _, ok := acic.mutation.ExpireDate(); !ok {
		return &ValidationError{Name: "expire_date", err: errors.New(`ent: missing required field "ActiveCodeInfo.expire_date"`)}
	}
	return nil
}

func (acic *ActiveCodeInfoCreate) sqlSave(ctx context.Context) (*ActiveCodeInfo, error) {
	if err := acic.check(); err != nil {
		return nil, err
	}
	_node, _spec := acic.createSpec()
	if err := sqlgraph.CreateNode(ctx, acic.driver, _spec); err != nil {
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
	acic.mutation.id = &_node.ID
	acic.mutation.done = true
	return _node, nil
}

func (acic *ActiveCodeInfoCreate) createSpec() (*ActiveCodeInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &ActiveCodeInfo{config: acic.config}
		_spec = sqlgraph.NewCreateSpec(activecodeinfo.Table, sqlgraph.NewFieldSpec(activecodeinfo.FieldID, field.TypeUUID))
	)
	if id, ok := acic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := acic.mutation.CreatedAt(); ok {
		_spec.SetField(activecodeinfo.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := acic.mutation.UpdatedAt(); ok {
		_spec.SetField(activecodeinfo.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := acic.mutation.Status(); ok {
		_spec.SetField(activecodeinfo.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := acic.mutation.ActiveKey(); ok {
		_spec.SetField(activecodeinfo.FieldActiveKey, field.TypeString, value)
		_node.ActiveKey = value
	}
	if value, ok := acic.mutation.UserID(); ok {
		_spec.SetField(activecodeinfo.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := acic.mutation.AppID(); ok {
		_spec.SetField(activecodeinfo.FieldAppID, field.TypeString, value)
		_node.AppID = value
	}
	if value, ok := acic.mutation.ActiveIP(); ok {
		_spec.SetField(activecodeinfo.FieldActiveIP, field.TypeString, value)
		_node.ActiveIP = value
	}
	if value, ok := acic.mutation.DeviceSn(); ok {
		_spec.SetField(activecodeinfo.FieldDeviceSn, field.TypeString, value)
		_node.DeviceSn = value
	}
	if value, ok := acic.mutation.DeviceMAC(); ok {
		_spec.SetField(activecodeinfo.FieldDeviceMAC, field.TypeString, value)
		_node.DeviceMAC = value
	}
	if value, ok := acic.mutation.DeviceIdentity(); ok {
		_spec.SetField(activecodeinfo.FieldDeviceIdentity, field.TypeString, value)
		_node.DeviceIdentity = value
	}
	if value, ok := acic.mutation.ActiveDate(); ok {
		_spec.SetField(activecodeinfo.FieldActiveDate, field.TypeTime, value)
		_node.ActiveDate = value
	}
	if value, ok := acic.mutation.ActiveType(); ok {
		_spec.SetField(activecodeinfo.FieldActiveType, field.TypeInt64, value)
		_node.ActiveType = value
	}
	if value, ok := acic.mutation.ActiveFile(); ok {
		_spec.SetField(activecodeinfo.FieldActiveFile, field.TypeString, value)
		_node.ActiveFile = value
	}
	if value, ok := acic.mutation.Version(); ok {
		_spec.SetField(activecodeinfo.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := acic.mutation.StartDate(); ok {
		_spec.SetField(activecodeinfo.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := acic.mutation.ExpireDate(); ok {
		_spec.SetField(activecodeinfo.FieldExpireDate, field.TypeTime, value)
		_node.ExpireDate = value
	}
	if nodes := acic.mutation.AppSdkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   activecodeinfo.AppSdkTable,
			Columns: []string{activecodeinfo.AppSdkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(appsdk.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AppSdkID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ActiveCodeInfoCreateBulk is the builder for creating many ActiveCodeInfo entities in bulk.
type ActiveCodeInfoCreateBulk struct {
	config
	err      error
	builders []*ActiveCodeInfoCreate
}

// Save creates the ActiveCodeInfo entities in the database.
func (acicb *ActiveCodeInfoCreateBulk) Save(ctx context.Context) ([]*ActiveCodeInfo, error) {
	if acicb.err != nil {
		return nil, acicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acicb.builders))
	nodes := make([]*ActiveCodeInfo, len(acicb.builders))
	mutators := make([]Mutator, len(acicb.builders))
	for i := range acicb.builders {
		func(i int, root context.Context) {
			builder := acicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActiveCodeInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, acicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acicb *ActiveCodeInfoCreateBulk) SaveX(ctx context.Context) []*ActiveCodeInfo {
	v, err := acicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acicb *ActiveCodeInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := acicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acicb *ActiveCodeInfoCreateBulk) ExecX(ctx context.Context) {
	if err := acicb.Exec(ctx); err != nil {
		panic(err)
	}
}
