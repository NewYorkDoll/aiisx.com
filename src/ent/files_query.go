// DO NOT EDIT, CODE GENERATED BY entc. yiziluoying

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"aiisx.com/src/ent/files"
	"aiisx.com/src/ent/post"
	"aiisx.com/src/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FilesQuery is the builder for querying Files entities.
type FilesQuery struct {
	config
	limit          *int
	offset         *int
	unique         *bool
	order          []OrderFunc
	fields         []string
	predicates     []predicate.Files
	withPosts      *PostQuery
	modifiers      []func(*sql.Selector)
	loadTotal      []func(context.Context, []*Files) error
	withNamedPosts map[string]*PostQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FilesQuery builder.
func (fq *FilesQuery) Where(ps ...predicate.Files) *FilesQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FilesQuery) Limit(limit int) *FilesQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FilesQuery) Offset(offset int) *FilesQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FilesQuery) Unique(unique bool) *FilesQuery {
	fq.unique = &unique
	return fq
}

// Order adds an order step to the query.
func (fq *FilesQuery) Order(o ...OrderFunc) *FilesQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryPosts chains the current query on the "posts" edge.
func (fq *FilesQuery) QueryPosts() *PostQuery {
	query := &PostQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(files.Table, files.FieldID, selector),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, files.PostsTable, files.PostsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Files entity from the query.
// Returns a *NotFoundError when no Files was found.
func (fq *FilesQuery) First(ctx context.Context) (*Files, error) {
	nodes, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{files.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FilesQuery) FirstX(ctx context.Context) *Files {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Files ID from the query.
// Returns a *NotFoundError when no Files ID was found.
func (fq *FilesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{files.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FilesQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Files entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Files entity is found.
// Returns a *NotFoundError when no Files entities are found.
func (fq *FilesQuery) Only(ctx context.Context) (*Files, error) {
	nodes, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{files.Label}
	default:
		return nil, &NotSingularError{files.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FilesQuery) OnlyX(ctx context.Context) *Files {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Files ID in the query.
// Returns a *NotSingularError when more than one Files ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FilesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{files.Label}
	default:
		err = &NotSingularError{files.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FilesQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FilesSlice.
func (fq *FilesQuery) All(ctx context.Context) ([]*Files, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FilesQuery) AllX(ctx context.Context) []*Files {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Files IDs.
func (fq *FilesQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := fq.Select(files.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FilesQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FilesQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FilesQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FilesQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FilesQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FilesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FilesQuery) Clone() *FilesQuery {
	if fq == nil {
		return nil
	}
	return &FilesQuery{
		config:     fq.config,
		limit:      fq.limit,
		offset:     fq.offset,
		order:      append([]OrderFunc{}, fq.order...),
		predicates: append([]predicate.Files{}, fq.predicates...),
		withPosts:  fq.withPosts.Clone(),
		// clone intermediate query.
		sql:    fq.sql.Clone(),
		path:   fq.path,
		unique: fq.unique,
	}
}

// WithPosts tells the query-builder to eager-load the nodes that are connected to
// the "posts" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FilesQuery) WithPosts(opts ...func(*PostQuery)) *FilesQuery {
	query := &PostQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withPosts = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Files.Query().
//		GroupBy(files.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FilesQuery) GroupBy(field string, fields ...string) *FilesGroupBy {
	grbuild := &FilesGroupBy{config: fq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(ctx), nil
	}
	grbuild.label = files.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Files.Query().
//		Select(files.FieldCreateTime).
//		Scan(ctx, &v)
func (fq *FilesQuery) Select(fields ...string) *FilesSelect {
	fq.fields = append(fq.fields, fields...)
	selbuild := &FilesSelect{FilesQuery: fq}
	selbuild.label = files.Label
	selbuild.flds, selbuild.scan = &fq.fields, selbuild.Scan
	return selbuild
}

func (fq *FilesQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fq.fields {
		if !files.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	if files.Policy == nil {
		return errors.New("ent: uninitialized files.Policy (forgotten import ent/runtime?)")
	}
	if err := files.Policy.EvalQuery(ctx, fq); err != nil {
		return err
	}
	return nil
}

func (fq *FilesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Files, error) {
	var (
		nodes       = []*Files{}
		_spec       = fq.querySpec()
		loadedTypes = [1]bool{
			fq.withPosts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Files).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Files{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withPosts; query != nil {
		if err := fq.loadPosts(ctx, query, nodes,
			func(n *Files) { n.Edges.Posts = []*Post{} },
			func(n *Files, e *Post) { n.Edges.Posts = append(n.Edges.Posts, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range fq.withNamedPosts {
		if err := fq.loadPosts(ctx, query, nodes,
			func(n *Files) { n.appendNamedPosts(name) },
			func(n *Files, e *Post) { n.appendNamedPosts(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range fq.loadTotal {
		if err := fq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FilesQuery) loadPosts(ctx context.Context, query *PostQuery, nodes []*Files, init func(*Files), assign func(*Files, *Post)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Files)
	nids := make(map[int]map[*Files]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(files.PostsTable)
		s.Join(joinT).On(s.C(post.FieldID), joinT.C(files.PostsPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(files.PostsPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(files.PostsPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]any, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]any{new(sql.NullInt64)}, values...), nil
		}
		spec.Assign = func(columns []string, values []any) error {
			outValue := int(values[0].(*sql.NullInt64).Int64)
			inValue := int(values[1].(*sql.NullInt64).Int64)
			if nids[inValue] == nil {
				nids[inValue] = map[*Files]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "posts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (fq *FilesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	_spec.Node.Columns = fq.fields
	if len(fq.fields) > 0 {
		_spec.Unique = fq.unique != nil && *fq.unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FilesQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := fq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (fq *FilesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   files.Table,
			Columns: files.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: files.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, files.FieldID)
		for i := range fields {
			if fields[i] != files.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FilesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(files.Table)
	columns := fq.fields
	if len(columns) == 0 {
		columns = files.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.unique != nil && *fq.unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedPosts tells the query-builder to eager-load the nodes that are connected to the "posts"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (fq *FilesQuery) WithNamedPosts(name string, opts ...func(*PostQuery)) *FilesQuery {
	query := &PostQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	if fq.withNamedPosts == nil {
		fq.withNamedPosts = make(map[string]*PostQuery)
	}
	fq.withNamedPosts[name] = query
	return fq
}

// FilesGroupBy is the group-by builder for Files entities.
type FilesGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FilesGroupBy) Aggregate(fns ...AggregateFunc) *FilesGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fgb *FilesGroupBy) Scan(ctx context.Context, v any) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

func (fgb *FilesGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range fgb.fields {
		if !files.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FilesGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql.Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
		for _, f := range fgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fgb.fields...)...)
}

// FilesSelect is the builder for selecting fields of Files entities.
type FilesSelect struct {
	*FilesQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FilesSelect) Scan(ctx context.Context, v any) error {
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	fs.sql = fs.FilesQuery.sqlQuery(ctx)
	return fs.sqlScan(ctx, v)
}

func (fs *FilesSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := fs.sql.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
