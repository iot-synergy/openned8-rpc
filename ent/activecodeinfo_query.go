// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent/activecodeinfo"
	"github.com/iot-synergy/openned8-rpc/ent/predicate"
)

// ActiveCodeInfoQuery is the builder for querying ActiveCodeInfo entities.
type ActiveCodeInfoQuery struct {
	config
	ctx        *QueryContext
	order      []activecodeinfo.OrderOption
	inters     []Interceptor
	predicates []predicate.ActiveCodeInfo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ActiveCodeInfoQuery builder.
func (aciq *ActiveCodeInfoQuery) Where(ps ...predicate.ActiveCodeInfo) *ActiveCodeInfoQuery {
	aciq.predicates = append(aciq.predicates, ps...)
	return aciq
}

// Limit the number of records to be returned by this query.
func (aciq *ActiveCodeInfoQuery) Limit(limit int) *ActiveCodeInfoQuery {
	aciq.ctx.Limit = &limit
	return aciq
}

// Offset to start from.
func (aciq *ActiveCodeInfoQuery) Offset(offset int) *ActiveCodeInfoQuery {
	aciq.ctx.Offset = &offset
	return aciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aciq *ActiveCodeInfoQuery) Unique(unique bool) *ActiveCodeInfoQuery {
	aciq.ctx.Unique = &unique
	return aciq
}

// Order specifies how the records should be ordered.
func (aciq *ActiveCodeInfoQuery) Order(o ...activecodeinfo.OrderOption) *ActiveCodeInfoQuery {
	aciq.order = append(aciq.order, o...)
	return aciq
}

// First returns the first ActiveCodeInfo entity from the query.
// Returns a *NotFoundError when no ActiveCodeInfo was found.
func (aciq *ActiveCodeInfoQuery) First(ctx context.Context) (*ActiveCodeInfo, error) {
	nodes, err := aciq.Limit(1).All(setContextOp(ctx, aciq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{activecodeinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aciq *ActiveCodeInfoQuery) FirstX(ctx context.Context) *ActiveCodeInfo {
	node, err := aciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ActiveCodeInfo ID from the query.
// Returns a *NotFoundError when no ActiveCodeInfo ID was found.
func (aciq *ActiveCodeInfoQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aciq.Limit(1).IDs(setContextOp(ctx, aciq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{activecodeinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aciq *ActiveCodeInfoQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ActiveCodeInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ActiveCodeInfo entity is found.
// Returns a *NotFoundError when no ActiveCodeInfo entities are found.
func (aciq *ActiveCodeInfoQuery) Only(ctx context.Context) (*ActiveCodeInfo, error) {
	nodes, err := aciq.Limit(2).All(setContextOp(ctx, aciq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{activecodeinfo.Label}
	default:
		return nil, &NotSingularError{activecodeinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aciq *ActiveCodeInfoQuery) OnlyX(ctx context.Context) *ActiveCodeInfo {
	node, err := aciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ActiveCodeInfo ID in the query.
// Returns a *NotSingularError when more than one ActiveCodeInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (aciq *ActiveCodeInfoQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aciq.Limit(2).IDs(setContextOp(ctx, aciq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{activecodeinfo.Label}
	default:
		err = &NotSingularError{activecodeinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aciq *ActiveCodeInfoQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ActiveCodeInfos.
func (aciq *ActiveCodeInfoQuery) All(ctx context.Context) ([]*ActiveCodeInfo, error) {
	ctx = setContextOp(ctx, aciq.ctx, "All")
	if err := aciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ActiveCodeInfo, *ActiveCodeInfoQuery]()
	return withInterceptors[[]*ActiveCodeInfo](ctx, aciq, qr, aciq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aciq *ActiveCodeInfoQuery) AllX(ctx context.Context) []*ActiveCodeInfo {
	nodes, err := aciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ActiveCodeInfo IDs.
func (aciq *ActiveCodeInfoQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if aciq.ctx.Unique == nil && aciq.path != nil {
		aciq.Unique(true)
	}
	ctx = setContextOp(ctx, aciq.ctx, "IDs")
	if err = aciq.Select(activecodeinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aciq *ActiveCodeInfoQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aciq *ActiveCodeInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aciq.ctx, "Count")
	if err := aciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aciq, querierCount[*ActiveCodeInfoQuery](), aciq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aciq *ActiveCodeInfoQuery) CountX(ctx context.Context) int {
	count, err := aciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aciq *ActiveCodeInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aciq.ctx, "Exist")
	switch _, err := aciq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aciq *ActiveCodeInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := aciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ActiveCodeInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aciq *ActiveCodeInfoQuery) Clone() *ActiveCodeInfoQuery {
	if aciq == nil {
		return nil
	}
	return &ActiveCodeInfoQuery{
		config:     aciq.config,
		ctx:        aciq.ctx.Clone(),
		order:      append([]activecodeinfo.OrderOption{}, aciq.order...),
		inters:     append([]Interceptor{}, aciq.inters...),
		predicates: append([]predicate.ActiveCodeInfo{}, aciq.predicates...),
		// clone intermediate query.
		sql:  aciq.sql.Clone(),
		path: aciq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ActiveCodeInfo.Query().
//		GroupBy(activecodeinfo.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aciq *ActiveCodeInfoQuery) GroupBy(field string, fields ...string) *ActiveCodeInfoGroupBy {
	aciq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ActiveCodeInfoGroupBy{build: aciq}
	grbuild.flds = &aciq.ctx.Fields
	grbuild.label = activecodeinfo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.ActiveCodeInfo.Query().
//		Select(activecodeinfo.FieldCreatedAt).
//		Scan(ctx, &v)
func (aciq *ActiveCodeInfoQuery) Select(fields ...string) *ActiveCodeInfoSelect {
	aciq.ctx.Fields = append(aciq.ctx.Fields, fields...)
	sbuild := &ActiveCodeInfoSelect{ActiveCodeInfoQuery: aciq}
	sbuild.label = activecodeinfo.Label
	sbuild.flds, sbuild.scan = &aciq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ActiveCodeInfoSelect configured with the given aggregations.
func (aciq *ActiveCodeInfoQuery) Aggregate(fns ...AggregateFunc) *ActiveCodeInfoSelect {
	return aciq.Select().Aggregate(fns...)
}

func (aciq *ActiveCodeInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aciq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aciq); err != nil {
				return err
			}
		}
	}
	for _, f := range aciq.ctx.Fields {
		if !activecodeinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aciq.path != nil {
		prev, err := aciq.path(ctx)
		if err != nil {
			return err
		}
		aciq.sql = prev
	}
	return nil
}

func (aciq *ActiveCodeInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ActiveCodeInfo, error) {
	var (
		nodes = []*ActiveCodeInfo{}
		_spec = aciq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ActiveCodeInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ActiveCodeInfo{config: aciq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aciq *ActiveCodeInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aciq.querySpec()
	_spec.Node.Columns = aciq.ctx.Fields
	if len(aciq.ctx.Fields) > 0 {
		_spec.Unique = aciq.ctx.Unique != nil && *aciq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aciq.driver, _spec)
}

func (aciq *ActiveCodeInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(activecodeinfo.Table, activecodeinfo.Columns, sqlgraph.NewFieldSpec(activecodeinfo.FieldID, field.TypeUUID))
	_spec.From = aciq.sql
	if unique := aciq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aciq.path != nil {
		_spec.Unique = true
	}
	if fields := aciq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, activecodeinfo.FieldID)
		for i := range fields {
			if fields[i] != activecodeinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aciq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aciq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aciq *ActiveCodeInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aciq.driver.Dialect())
	t1 := builder.Table(activecodeinfo.Table)
	columns := aciq.ctx.Fields
	if len(columns) == 0 {
		columns = activecodeinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aciq.sql != nil {
		selector = aciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aciq.ctx.Unique != nil && *aciq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aciq.predicates {
		p(selector)
	}
	for _, p := range aciq.order {
		p(selector)
	}
	if offset := aciq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aciq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ActiveCodeInfoGroupBy is the group-by builder for ActiveCodeInfo entities.
type ActiveCodeInfoGroupBy struct {
	selector
	build *ActiveCodeInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acigb *ActiveCodeInfoGroupBy) Aggregate(fns ...AggregateFunc) *ActiveCodeInfoGroupBy {
	acigb.fns = append(acigb.fns, fns...)
	return acigb
}

// Scan applies the selector query and scans the result into the given value.
func (acigb *ActiveCodeInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acigb.build.ctx, "GroupBy")
	if err := acigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ActiveCodeInfoQuery, *ActiveCodeInfoGroupBy](ctx, acigb.build, acigb, acigb.build.inters, v)
}

func (acigb *ActiveCodeInfoGroupBy) sqlScan(ctx context.Context, root *ActiveCodeInfoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(acigb.fns))
	for _, fn := range acigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*acigb.flds)+len(acigb.fns))
		for _, f := range *acigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*acigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ActiveCodeInfoSelect is the builder for selecting fields of ActiveCodeInfo entities.
type ActiveCodeInfoSelect struct {
	*ActiveCodeInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (acis *ActiveCodeInfoSelect) Aggregate(fns ...AggregateFunc) *ActiveCodeInfoSelect {
	acis.fns = append(acis.fns, fns...)
	return acis
}

// Scan applies the selector query and scans the result into the given value.
func (acis *ActiveCodeInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acis.ctx, "Select")
	if err := acis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ActiveCodeInfoQuery, *ActiveCodeInfoSelect](ctx, acis.ActiveCodeInfoQuery, acis, acis.inters, v)
}

func (acis *ActiveCodeInfoSelect) sqlScan(ctx context.Context, root *ActiveCodeInfoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(acis.fns))
	for _, fn := range acis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*acis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
