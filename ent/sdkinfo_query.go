// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	uuid "github.com/gofrs/uuid/v5"
	"github.com/iot-synergy/openned8-rpc/ent/appsdk"
	"github.com/iot-synergy/openned8-rpc/ent/predicate"
	"github.com/iot-synergy/openned8-rpc/ent/sdkinfo"
)

// SdkInfoQuery is the builder for querying SdkInfo entities.
type SdkInfoQuery struct {
	config
	ctx        *QueryContext
	order      []sdkinfo.OrderOption
	inters     []Interceptor
	predicates []predicate.SdkInfo
	withAppSdk *AppSdkQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SdkInfoQuery builder.
func (siq *SdkInfoQuery) Where(ps ...predicate.SdkInfo) *SdkInfoQuery {
	siq.predicates = append(siq.predicates, ps...)
	return siq
}

// Limit the number of records to be returned by this query.
func (siq *SdkInfoQuery) Limit(limit int) *SdkInfoQuery {
	siq.ctx.Limit = &limit
	return siq
}

// Offset to start from.
func (siq *SdkInfoQuery) Offset(offset int) *SdkInfoQuery {
	siq.ctx.Offset = &offset
	return siq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (siq *SdkInfoQuery) Unique(unique bool) *SdkInfoQuery {
	siq.ctx.Unique = &unique
	return siq
}

// Order specifies how the records should be ordered.
func (siq *SdkInfoQuery) Order(o ...sdkinfo.OrderOption) *SdkInfoQuery {
	siq.order = append(siq.order, o...)
	return siq
}

// QueryAppSdk chains the current query on the "app_sdk" edge.
func (siq *SdkInfoQuery) QueryAppSdk() *AppSdkQuery {
	query := (&AppSdkClient{config: siq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := siq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := siq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sdkinfo.Table, sdkinfo.FieldID, selector),
			sqlgraph.To(appsdk.Table, appsdk.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sdkinfo.AppSdkTable, sdkinfo.AppSdkColumn),
		)
		fromU = sqlgraph.SetNeighbors(siq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SdkInfo entity from the query.
// Returns a *NotFoundError when no SdkInfo was found.
func (siq *SdkInfoQuery) First(ctx context.Context) (*SdkInfo, error) {
	nodes, err := siq.Limit(1).All(setContextOp(ctx, siq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sdkinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (siq *SdkInfoQuery) FirstX(ctx context.Context) *SdkInfo {
	node, err := siq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SdkInfo ID from the query.
// Returns a *NotFoundError when no SdkInfo ID was found.
func (siq *SdkInfoQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = siq.Limit(1).IDs(setContextOp(ctx, siq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sdkinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (siq *SdkInfoQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := siq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SdkInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SdkInfo entity is found.
// Returns a *NotFoundError when no SdkInfo entities are found.
func (siq *SdkInfoQuery) Only(ctx context.Context) (*SdkInfo, error) {
	nodes, err := siq.Limit(2).All(setContextOp(ctx, siq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sdkinfo.Label}
	default:
		return nil, &NotSingularError{sdkinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (siq *SdkInfoQuery) OnlyX(ctx context.Context) *SdkInfo {
	node, err := siq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SdkInfo ID in the query.
// Returns a *NotSingularError when more than one SdkInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (siq *SdkInfoQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = siq.Limit(2).IDs(setContextOp(ctx, siq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sdkinfo.Label}
	default:
		err = &NotSingularError{sdkinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (siq *SdkInfoQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := siq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SdkInfos.
func (siq *SdkInfoQuery) All(ctx context.Context) ([]*SdkInfo, error) {
	ctx = setContextOp(ctx, siq.ctx, "All")
	if err := siq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SdkInfo, *SdkInfoQuery]()
	return withInterceptors[[]*SdkInfo](ctx, siq, qr, siq.inters)
}

// AllX is like All, but panics if an error occurs.
func (siq *SdkInfoQuery) AllX(ctx context.Context) []*SdkInfo {
	nodes, err := siq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SdkInfo IDs.
func (siq *SdkInfoQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if siq.ctx.Unique == nil && siq.path != nil {
		siq.Unique(true)
	}
	ctx = setContextOp(ctx, siq.ctx, "IDs")
	if err = siq.Select(sdkinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (siq *SdkInfoQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := siq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (siq *SdkInfoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, siq.ctx, "Count")
	if err := siq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, siq, querierCount[*SdkInfoQuery](), siq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (siq *SdkInfoQuery) CountX(ctx context.Context) int {
	count, err := siq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (siq *SdkInfoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, siq.ctx, "Exist")
	switch _, err := siq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (siq *SdkInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := siq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SdkInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (siq *SdkInfoQuery) Clone() *SdkInfoQuery {
	if siq == nil {
		return nil
	}
	return &SdkInfoQuery{
		config:     siq.config,
		ctx:        siq.ctx.Clone(),
		order:      append([]sdkinfo.OrderOption{}, siq.order...),
		inters:     append([]Interceptor{}, siq.inters...),
		predicates: append([]predicate.SdkInfo{}, siq.predicates...),
		withAppSdk: siq.withAppSdk.Clone(),
		// clone intermediate query.
		sql:  siq.sql.Clone(),
		path: siq.path,
	}
}

// WithAppSdk tells the query-builder to eager-load the nodes that are connected to
// the "app_sdk" edge. The optional arguments are used to configure the query builder of the edge.
func (siq *SdkInfoQuery) WithAppSdk(opts ...func(*AppSdkQuery)) *SdkInfoQuery {
	query := (&AppSdkClient{config: siq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	siq.withAppSdk = query
	return siq
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
//	client.SdkInfo.Query().
//		GroupBy(sdkinfo.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (siq *SdkInfoQuery) GroupBy(field string, fields ...string) *SdkInfoGroupBy {
	siq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SdkInfoGroupBy{build: siq}
	grbuild.flds = &siq.ctx.Fields
	grbuild.label = sdkinfo.Label
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
//	client.SdkInfo.Query().
//		Select(sdkinfo.FieldCreatedAt).
//		Scan(ctx, &v)
func (siq *SdkInfoQuery) Select(fields ...string) *SdkInfoSelect {
	siq.ctx.Fields = append(siq.ctx.Fields, fields...)
	sbuild := &SdkInfoSelect{SdkInfoQuery: siq}
	sbuild.label = sdkinfo.Label
	sbuild.flds, sbuild.scan = &siq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SdkInfoSelect configured with the given aggregations.
func (siq *SdkInfoQuery) Aggregate(fns ...AggregateFunc) *SdkInfoSelect {
	return siq.Select().Aggregate(fns...)
}

func (siq *SdkInfoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range siq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, siq); err != nil {
				return err
			}
		}
	}
	for _, f := range siq.ctx.Fields {
		if !sdkinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if siq.path != nil {
		prev, err := siq.path(ctx)
		if err != nil {
			return err
		}
		siq.sql = prev
	}
	return nil
}

func (siq *SdkInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SdkInfo, error) {
	var (
		nodes       = []*SdkInfo{}
		_spec       = siq.querySpec()
		loadedTypes = [1]bool{
			siq.withAppSdk != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SdkInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SdkInfo{config: siq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, siq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := siq.withAppSdk; query != nil {
		if err := siq.loadAppSdk(ctx, query, nodes,
			func(n *SdkInfo) { n.Edges.AppSdk = []*AppSdk{} },
			func(n *SdkInfo, e *AppSdk) { n.Edges.AppSdk = append(n.Edges.AppSdk, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (siq *SdkInfoQuery) loadAppSdk(ctx context.Context, query *AppSdkQuery, nodes []*SdkInfo, init func(*SdkInfo), assign func(*SdkInfo, *AppSdk)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*SdkInfo)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(appsdk.FieldSdk)
	}
	query.Where(predicate.AppSdk(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(sdkinfo.AppSdkColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.Sdk
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "sdk" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (siq *SdkInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := siq.querySpec()
	_spec.Node.Columns = siq.ctx.Fields
	if len(siq.ctx.Fields) > 0 {
		_spec.Unique = siq.ctx.Unique != nil && *siq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, siq.driver, _spec)
}

func (siq *SdkInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(sdkinfo.Table, sdkinfo.Columns, sqlgraph.NewFieldSpec(sdkinfo.FieldID, field.TypeUUID))
	_spec.From = siq.sql
	if unique := siq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if siq.path != nil {
		_spec.Unique = true
	}
	if fields := siq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sdkinfo.FieldID)
		for i := range fields {
			if fields[i] != sdkinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := siq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := siq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := siq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := siq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (siq *SdkInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(siq.driver.Dialect())
	t1 := builder.Table(sdkinfo.Table)
	columns := siq.ctx.Fields
	if len(columns) == 0 {
		columns = sdkinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if siq.sql != nil {
		selector = siq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if siq.ctx.Unique != nil && *siq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range siq.predicates {
		p(selector)
	}
	for _, p := range siq.order {
		p(selector)
	}
	if offset := siq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := siq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SdkInfoGroupBy is the group-by builder for SdkInfo entities.
type SdkInfoGroupBy struct {
	selector
	build *SdkInfoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sigb *SdkInfoGroupBy) Aggregate(fns ...AggregateFunc) *SdkInfoGroupBy {
	sigb.fns = append(sigb.fns, fns...)
	return sigb
}

// Scan applies the selector query and scans the result into the given value.
func (sigb *SdkInfoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sigb.build.ctx, "GroupBy")
	if err := sigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SdkInfoQuery, *SdkInfoGroupBy](ctx, sigb.build, sigb, sigb.build.inters, v)
}

func (sigb *SdkInfoGroupBy) sqlScan(ctx context.Context, root *SdkInfoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sigb.fns))
	for _, fn := range sigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sigb.flds)+len(sigb.fns))
		for _, f := range *sigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SdkInfoSelect is the builder for selecting fields of SdkInfo entities.
type SdkInfoSelect struct {
	*SdkInfoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sis *SdkInfoSelect) Aggregate(fns ...AggregateFunc) *SdkInfoSelect {
	sis.fns = append(sis.fns, fns...)
	return sis
}

// Scan applies the selector query and scans the result into the given value.
func (sis *SdkInfoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sis.ctx, "Select")
	if err := sis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SdkInfoQuery, *SdkInfoSelect](ctx, sis.SdkInfoQuery, sis, sis.inters, v)
}

func (sis *SdkInfoSelect) sqlScan(ctx context.Context, root *SdkInfoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sis.fns))
	for _, fn := range sis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
