// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/iot-synergy/openned8-rpc/ent/categoryinfo"
	"github.com/iot-synergy/openned8-rpc/ent/predicate"
)

// CategoryInfoQuery is the builder for querying CategoryInfo entities.
type CategoryInfoQuery struct {
	config
	ctx        *QueryContext
	order      []categoryinfo.OrderOption
	inters     []Interceptor
	predicates []predicate.CategoryInfo
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CategoryInfoQuery builder.
func (ciq *CategoryInfoQuery) Where(ps ...predicate.CategoryInfo) *CategoryInfoQuery {
	ciq.predicates = append(ciq.predicates, ps...)
	return ciq
}

// Limit the number of records to be returned by this query.
func (ciq *CategoryInfoQuery) Limit(limit int) *CategoryInfoQuery {
	ciq.ctx.Limit = &limit
	return ciq
}

// Offset to start from.
func (ciq *CategoryInfoQuery) Offset(offset int) *CategoryInfoQuery {
	ciq.ctx.Offset = &offset
	return ciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ciq *CategoryInfoQuery) Unique(unique bool) *CategoryInfoQuery {
	ciq.ctx.Unique = &unique
	return ciq
}

// Order specifies how the records should be ordered.
func (ciq *CategoryInfoQuery) Order(o ...categoryinfo.OrderOption) *CategoryInfoQuery {
	ciq.order = append(ciq.order, o...)
	return ciq
}

// First returns the first CategoryInfo entity from the query.
// Returns a *NotFoundError when no CategoryInfo was found.
func (ciq *CategoryInfoQuery) First(ctx context.Context) (*CategoryInfo, error) {
	nodes, err := ciq.Limit(1).All(setContextOp(ctx, ciq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{categoryinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ciq *CategoryInfoQuery) FirstX(ctx context.Context) *CategoryInfo {
	node, err := ciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CategoryInfo ID from the query.
// Returns a *NotFoundError when no CategoryInfo ID was found.
func (ciq *CategoryInfoQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ciq.Limit(1).IDs(setContextOp(ctx, ciq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{categoryinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ciq *CategoryInfoQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := ciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CategoryInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CategoryInfo entity is found.
// Returns a *NotFoundError when no CategoryInfo entities are found.
func (ciq *CategoryInfoQuery) Only(ctx context.Context) (*CategoryInfo, error) {
	nodes, err := ciq.Limit(2).All(setContextOp(ctx, ciq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{categoryinfo.Label}
	default:
		return nil, &NotSingularError{categoryinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ciq *CategoryInfoQuery) OnlyX(ctx context.Context) *CategoryInfo {
	node, err := ciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CategoryInfo ID in the query.
// Returns a *NotSingularError when more than one CategoryInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (ciq *CategoryInfoQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ciq.Limit(2).IDs(setContextOp(ctx, ciq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{categoryinfo.Label}
	default:
		err = &NotSingularError{categoryinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ciq *CategoryInfoQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := ciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CategoryInfos.
func (ciq *CategoryInfoQuery) All(ctx context.Context) ([]*CategoryInfo, error) {
	ctx = setContextOp(ctx, ciq.ctx, "All")
	if err := ciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CategoryInfo, *CategoryInfoQuery]()
	return withInterceptors[[]*CategoryInfo](ctx, ciq, qr, ciq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ciq *CategoryInfoQuery) AllX(ctx context.Context) []*CategoryInfo {
	nodes, err := ciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CategoryInfo IDs.
func (ciq *CategoryInfoQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if ciq.ctx.Unique == nil && ciq.path != nil {
		ciq.Unique(true)
	}
	ctx = setContextOp(ctx, ciq.ctx, "IDs")
	if err = ciq.Select(categoryinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ciq *CategoryInfoQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := ciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ciq *CategoryInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ciq.ctx, "Count")
	if err := ciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ciq, querierCount[*CategoryInfoQuery](), ciq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ciq *CategoryInfoQuery) CountX(ctx context.Context) int {
	count, err := ciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ciq *CategoryInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ciq.ctx, "Exist")
	switch _, err := ciq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ciq *CategoryInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := ciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CategoryInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ciq *CategoryInfoQuery) Clone() *CategoryInfoQuery {
	if ciq == nil {
		return nil
	}
	return &CategoryInfoQuery{
		config:     ciq.config,
		ctx:        ciq.ctx.Clone(),
		order:      append([]categoryinfo.OrderOption{}, ciq.order...),
		inters:     append([]Interceptor{}, ciq.inters...),
		predicates: append([]predicate.CategoryInfo{}, ciq.predicates...),
		// clone intermediate query.
		sql:  ciq.sql.Clone(),
		path: ciq.path,
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
//	client.CategoryInfo.Query().
//		GroupBy(categoryinfo.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ciq *CategoryInfoQuery) GroupBy(field string, fields ...string) *CategoryInfoGroupBy {
	ciq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CategoryInfoGroupBy{build: ciq}
	grbuild.flds = &ciq.ctx.Fields
	grbuild.label = categoryinfo.Label
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
//	client.CategoryInfo.Query().
//		Select(categoryinfo.FieldCreatedAt).
//		Scan(ctx, &v)
func (ciq *CategoryInfoQuery) Select(fields ...string) *CategoryInfoSelect {
	ciq.ctx.Fields = append(ciq.ctx.Fields, fields...)
	sbuild := &CategoryInfoSelect{CategoryInfoQuery: ciq}
	sbuild.label = categoryinfo.Label
	sbuild.flds, sbuild.scan = &ciq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CategoryInfoSelect configured with the given aggregations.
func (ciq *CategoryInfoQuery) Aggregate(fns ...AggregateFunc) *CategoryInfoSelect {
	return ciq.Select().Aggregate(fns...)
}

func (ciq *CategoryInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ciq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ciq); err != nil {
				return err
			}
		}
	}
	for _, f := range ciq.ctx.Fields {
		if !categoryinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ciq.path != nil {
		prev, err := ciq.path(ctx)
		if err != nil {
			return err
		}
		ciq.sql = prev
	}
	return nil
}

func (ciq *CategoryInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CategoryInfo, error) {
	var (
		nodes = []*CategoryInfo{}
		_spec = ciq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CategoryInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CategoryInfo{config: ciq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ciq *CategoryInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ciq.querySpec()
	_spec.Node.Columns = ciq.ctx.Fields
	if len(ciq.ctx.Fields) > 0 {
		_spec.Unique = ciq.ctx.Unique != nil && *ciq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ciq.driver, _spec)
}

func (ciq *CategoryInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(categoryinfo.Table, categoryinfo.Columns, sqlgraph.NewFieldSpec(categoryinfo.FieldID, field.TypeUint64))
	_spec.From = ciq.sql
	if unique := ciq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ciq.path != nil {
		_spec.Unique = true
	}
	if fields := ciq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, categoryinfo.FieldID)
		for i := range fields {
			if fields[i] != categoryinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ciq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ciq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ciq *CategoryInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ciq.driver.Dialect())
	t1 := builder.Table(categoryinfo.Table)
	columns := ciq.ctx.Fields
	if len(columns) == 0 {
		columns = categoryinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ciq.sql != nil {
		selector = ciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ciq.ctx.Unique != nil && *ciq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ciq.predicates {
		p(selector)
	}
	for _, p := range ciq.order {
		p(selector)
	}
	if offset := ciq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ciq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CategoryInfoGroupBy is the group-by builder for CategoryInfo entities.
type CategoryInfoGroupBy struct {
	selector
	build *CategoryInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cigb *CategoryInfoGroupBy) Aggregate(fns ...AggregateFunc) *CategoryInfoGroupBy {
	cigb.fns = append(cigb.fns, fns...)
	return cigb
}

// Scan applies the selector query and scans the result into the given value.
func (cigb *CategoryInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cigb.build.ctx, "GroupBy")
	if err := cigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CategoryInfoQuery, *CategoryInfoGroupBy](ctx, cigb.build, cigb, cigb.build.inters, v)
}

func (cigb *CategoryInfoGroupBy) sqlScan(ctx context.Context, root *CategoryInfoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cigb.fns))
	for _, fn := range cigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cigb.flds)+len(cigb.fns))
		for _, f := range *cigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CategoryInfoSelect is the builder for selecting fields of CategoryInfo entities.
type CategoryInfoSelect struct {
	*CategoryInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cis *CategoryInfoSelect) Aggregate(fns ...AggregateFunc) *CategoryInfoSelect {
	cis.fns = append(cis.fns, fns...)
	return cis
}

// Scan applies the selector query and scans the result into the given value.
func (cis *CategoryInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cis.ctx, "Select")
	if err := cis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CategoryInfoQuery, *CategoryInfoSelect](ctx, cis.CategoryInfoQuery, cis, cis.inters, v)
}

func (cis *CategoryInfoSelect) sqlScan(ctx context.Context, root *CategoryInfoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cis.fns))
	for _, fn := range cis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
