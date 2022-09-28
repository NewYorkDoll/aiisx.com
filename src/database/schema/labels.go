package schema

import (
	"context"
	"regexp"

	"aiisx.com/src/ent/privacy"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var reLabel = regexp.MustCompile(`^[0-9a-zA-Z\d][0-9a-zA-Z\d-]*$`)

type Label struct {
	ent.Schema
}

func (Label) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Match(reLabel).Annotations(
			entgql.OrderField("NAME"),
		),
	}
}

func (Label) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Label) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.ContextQueryMutationRule(func(ctx context.Context) error { return privacy.Allow }),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

func (Label) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).Annotations(
			entgql.RelayConnection(),
		),
		edge.To("github_repositories", GithubRepository.Type).Annotations(
			entgql.RelayConnection(),
		),
	}
}

func (Label) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
