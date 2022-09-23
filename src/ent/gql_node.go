// DO NOT EDIT, CODE GENERATED BY entc. yiziluoying

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"sync/atomic"

	"aiisx.com/src/ent/githubevent"
	"aiisx.com/src/ent/githubrepository"
	"aiisx.com/src/ent/label"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"golang.org/x/sync/semaphore"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     int      `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string `json:"type,omitempty"` // edge type.
	Name string `json:"name,omitempty"` // edge name.
	IDs  []int  `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (ge *GithubEvent) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ge.ID,
		Type:   "GithubEvent",
		Fields: make([]*Field, 9),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(ge.EventID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "event_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ge.EventType); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "event_type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ge.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ge.Public); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "bool",
		Name:  "public",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ge.ActorID); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "int64",
		Name:  "actor_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ge.Actor); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "*github.User",
		Name:  "actor",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ge.RepoID); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "int64",
		Name:  "repo_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ge.Repo); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "*github.Repository",
		Name:  "repo",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ge.Payload); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "map[string]interface {}",
		Name:  "payload",
		Value: string(buf),
	}
	return node, nil
}

func (gr *GithubRepository) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     gr.ID,
		Type:   "GithubRepository",
		Fields: make([]*Field, 19),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(gr.RepoID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int64",
		Name:  "repo_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.FullName); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "full_name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.OwnerLogin); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "owner_login",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.Owner); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "*github.User",
		Name:  "owner",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.Public); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "bool",
		Name:  "public",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.HTMLURL); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "html_url",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.Description); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.Fork); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "bool",
		Name:  "fork",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.Homepage); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "string",
		Name:  "homepage",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.StarCount); err != nil {
		return nil, err
	}
	node.Fields[10] = &Field{
		Type:  "int",
		Name:  "star_count",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.DefaultBranch); err != nil {
		return nil, err
	}
	node.Fields[11] = &Field{
		Type:  "string",
		Name:  "default_branch",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.IsTemplate); err != nil {
		return nil, err
	}
	node.Fields[12] = &Field{
		Type:  "bool",
		Name:  "is_template",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.HasIssues); err != nil {
		return nil, err
	}
	node.Fields[13] = &Field{
		Type:  "bool",
		Name:  "has_issues",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.Archived); err != nil {
		return nil, err
	}
	node.Fields[14] = &Field{
		Type:  "bool",
		Name:  "archived",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.PushedAt); err != nil {
		return nil, err
	}
	node.Fields[15] = &Field{
		Type:  "time.Time",
		Name:  "pushed_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[16] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[17] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.License); err != nil {
		return nil, err
	}
	node.Fields[18] = &Field{
		Type:  "*github.License",
		Name:  "license",
		Value: string(buf),
	}
	return node, nil
}

func (l *Label) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     l.ID,
		Type:   "Label",
		Fields: make([]*Field, 0),
		Edges:  make([]*Edge, 0),
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id int) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, int) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, int) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, int) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id int) (string, error) {
			return c.tables.nodeType(ctx, c.driver, id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id int, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id int) (Noder, error) {
	switch table {
	case githubevent.Table:
		query := c.GithubEvent.Query().
			Where(githubevent.ID(id))
		query, err := query.CollectFields(ctx, "GithubEvent")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case githubrepository.Table:
		query := c.GithubRepository.Query().
			Where(githubrepository.ID(id))
		query, err := query.CollectFields(ctx, "GithubRepository")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case label.Table:
		query := c.Label.Query().
			Where(label.ID(id))
		query, err := query.CollectFields(ctx, "Label")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []int, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]int)
	id2idx := make(map[int][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []int) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[int][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case githubevent.Table:
		query := c.GithubEvent.Query().
			Where(githubevent.IDIn(ids...))
		query, err := query.CollectFields(ctx, "GithubEvent")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case githubrepository.Table:
		query := c.GithubRepository.Query().
			Where(githubrepository.IDIn(ids...))
		query, err := query.CollectFields(ctx, "GithubRepository")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case label.Table:
		query := c.Label.Query().
			Where(label.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Label")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

type tables struct {
	once  sync.Once
	sem   *semaphore.Weighted
	value atomic.Value
}

func (t *tables) nodeType(ctx context.Context, drv dialect.Driver, id int) (string, error) {
	tables, err := t.Load(ctx, drv)
	if err != nil {
		return "", err
	}
	idx := int(id / (1<<32 - 1))
	if idx < 0 || idx >= len(tables) {
		return "", fmt.Errorf("cannot resolve table from id %v: %w", id, errNodeInvalidID)
	}
	return tables[idx], nil
}

func (t *tables) Load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	t.once.Do(func() { t.sem = semaphore.NewWeighted(1) })
	if err := t.sem.Acquire(ctx, 1); err != nil {
		return nil, err
	}
	defer t.sem.Release(1)
	if tables := t.value.Load(); tables != nil {
		return tables.([]string), nil
	}
	tables, err := t.load(ctx, drv)
	if err == nil {
		t.value.Store(tables)
	}
	return tables, err
}

func (*tables) load(ctx context.Context, drv dialect.Driver) ([]string, error) {
	rows := &sql.Rows{}
	query, args := sql.Dialect(drv.Dialect()).
		Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := drv.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var tables []string
	return tables, sql.ScanSlice(rows, &tables)
}
