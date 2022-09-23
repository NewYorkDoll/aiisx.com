// DO NOT EDIT, CODE GENERATED BY entc. yiziluoying

package ent

import (
	"context"
	"fmt"

	"aiisx.com/src/ent/githubevent"
	"aiisx.com/src/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GithubEventDelete is the builder for deleting a GithubEvent entity.
type GithubEventDelete struct {
	config
	hooks    []Hook
	mutation *GithubEventMutation
}

// Where appends a list predicates to the GithubEventDelete builder.
func (ged *GithubEventDelete) Where(ps ...predicate.GithubEvent) *GithubEventDelete {
	ged.mutation.Where(ps...)
	return ged
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ged *GithubEventDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ged.hooks) == 0 {
		affected, err = ged.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GithubEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ged.mutation = mutation
			affected, err = ged.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ged.hooks) - 1; i >= 0; i-- {
			if ged.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ged.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ged.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ged *GithubEventDelete) ExecX(ctx context.Context) int {
	n, err := ged.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ged *GithubEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: githubevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: githubevent.FieldID,
			},
		},
	}
	if ps := ged.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ged.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// GithubEventDeleteOne is the builder for deleting a single GithubEvent entity.
type GithubEventDeleteOne struct {
	ged *GithubEventDelete
}

// Exec executes the deletion query.
func (gedo *GithubEventDeleteOne) Exec(ctx context.Context) error {
	n, err := gedo.ged.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{githubevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gedo *GithubEventDeleteOne) ExecX(ctx context.Context) {
	gedo.ged.ExecX(ctx)
}
