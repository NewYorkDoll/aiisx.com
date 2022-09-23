// DO NOT EDIT, CODE GENERATED BY entc. yiziluoying

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ge *GithubEventQuery) CollectFields(ctx context.Context, satisfies ...string) (*GithubEventQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ge, nil
	}
	if err := ge.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ge, nil
}

func (ge *GithubEventQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	return nil
}

type githubeventPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GithubEventPaginateOption
}

func newGithubEventPaginateArgs(rv map[string]interface{}) *githubeventPaginateArgs {
	args := &githubeventPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &GithubEventOrder{Field: &GithubEventOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithGithubEventOrder(order))
			}
		case *GithubEventOrder:
			if v != nil {
				args.opts = append(args.opts, WithGithubEventOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*GithubEventWhereInput); ok {
		args.opts = append(args.opts, WithGithubEventFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gr *GithubRepositoryQuery) CollectFields(ctx context.Context, satisfies ...string) (*GithubRepositoryQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return gr, nil
	}
	if err := gr.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return gr, nil
}

func (gr *GithubRepositoryQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	return nil
}

type githubrepositoryPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []GithubRepositoryPaginateOption
}

func newGithubRepositoryPaginateArgs(rv map[string]interface{}) *githubrepositoryPaginateArgs {
	args := &githubrepositoryPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &GithubRepositoryOrder{Field: &GithubRepositoryOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithGithubRepositoryOrder(order))
			}
		case *GithubRepositoryOrder:
			if v != nil {
				args.opts = append(args.opts, WithGithubRepositoryOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*GithubRepositoryWhereInput); ok {
		args.opts = append(args.opts, WithGithubRepositoryFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (l *LabelQuery) CollectFields(ctx context.Context, satisfies ...string) (*LabelQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return l, nil
	}
	if err := l.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return l, nil
}

func (l *LabelQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	return nil
}

type labelPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []LabelPaginateOption
}

func newLabelPaginateArgs(rv map[string]interface{}) *labelPaginateArgs {
	args := &labelPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*LabelWhereInput); ok {
		args.opts = append(args.opts, WithLabelFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}
