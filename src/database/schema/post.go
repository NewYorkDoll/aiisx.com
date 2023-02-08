package schema

import (
	"context"
	"regexp"
	"time"

	"aiisx.com/src/ent/privacy"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var rePostSlug = regexp.MustCompile(`(?i)^[a-z\d][a-z\d-]{4,50}$`)

type Post struct {
	ent.Schema
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").Match(rePostSlug).Unique().Annotations(
			entgql.OrderField("SLUG"),
		),
		field.String("title").MaxLen(100).Annotations(
			entgql.OrderField("TITLE"),
		),
		field.String("content").SchemaType(map[string]string{
			dialect.MySQL: "text", // Override MySQL.
		}).MaxLen(15000).NotEmpty(),
		field.String("content_html").SchemaType(map[string]string{
			dialect.MySQL: "text", // Override MySQL.
		}).MaxLen(50000).NotEmpty().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
		),
		field.String("summary").SchemaType(map[string]string{
			dialect.MySQL: "text", // Override MySQL.
		}).MaxLen(10000).NotEmpty().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
		),
		field.Time("published_at").Default(time.Now).Annotations(
			entgql.OrderField("DATE"),
		),
		field.Int("view_count").Default(0).NonNegative().Annotations(
			entgql.OrderField("VIEW_COUNT"),
			entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
		),
		field.Bool("public").Default(false),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Post) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.ContextQueryMutationRule(func(ctx context.Context) error { return privacy.Allow }),
		},
		Query: privacy.QueryPolicy{

			privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
				type PublicFilter interface {
					WherePublic(p entql.BoolP)
				}
				_, ok := f.(PublicFilter)
				if !ok {
					return privacy.Denyf("missing public field in filter")
				}
				return privacy.Skip
			}),
			privacy.AlwaysAllowRule(),
		},
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("posts").Unique().Required().Annotations(
			entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput),
		),
		edge.From("labels", Label.Type).Ref("posts").Annotations(
			entgql.RelayConnection(),
		),
		edge.From("files", Files.Type).Ref("posts").Annotations(
			entgql.RelayConnection(),
		),
	}
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
