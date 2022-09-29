package database

import (
	"context"
	"strings"

	"aiisx.com/src/ent"
	"aiisx.com/src/ent/hook"
	"aiisx.com/src/ent/post"
	"aiisx.com/src/markdown"
	"github.com/apex/log"
	"github.com/microcosm-cc/bluemonday"
)

const summaryLen = 40

var htmlPolicy = bluemonday.StrictPolicy()

// RegisterHooks initializes all runtime hooks into the database client.
func RegisterHooks(ctx context.Context) {
	log.FromContext(ctx).Info("registering database hooks")

	db := ent.FromContext(ctx)
	if db == nil {
		panic("database client is nil")
	}

	db.Post.Use(
		hook.On(
			hook.If(func(next ent.Mutator) ent.Mutator {
				return hook.PostFunc(func(ctx context.Context, m *ent.PostMutation) (ent.Value, error) {
					content, ok := m.Content()
					if !ok {
						return next.Mutate(ctx, m)
					}

					md, err := markdown.Generate(ctx, content)
					if err != nil {
						return m, err
					}

					m.SetContentHTML(md)

					summary := strings.Fields(htmlPolicy.Sanitize(md))
					if len(summary) > summaryLen {
						m.SetSummary(strings.Join(summary[:summaryLen], " ") + "...")
					} else {
						m.SetSummary(strings.Join(summary, " ") + "...")
					}

					return next.Mutate(ctx, m)
				})
			}, hook.HasFields(post.FieldContent)),
			ent.OpCreate|ent.OpUpdateOne|ent.OpUpdate,
		),
	)
}
