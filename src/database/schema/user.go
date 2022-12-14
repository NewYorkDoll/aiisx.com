package schema

import (
	"regexp"

	"aiisx.com/src/ent/privacy"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var (
	reUserLogin = regexp.MustCompile(`(?i)^[a-z\d][a-z\d-]{0,38}$`)
)

type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Unique().Positive().Immutable(),
		field.String("login").Unique().Match(reUserLogin).Annotations(
			entgql.OrderField("LOGIN"),
		),
		field.String("name").Optional().MaxLen(400).Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("avatar_url").Optional().MaxLen(2048),
		field.String("html_url").Optional().MaxLen(2048),
		field.String("email").Optional().MaxLen(320).Annotations(
			entgql.OrderField("EMAIL"),
		),
		field.String("location").Optional().MaxLen(400),
		field.String("bio").Optional().MaxLen(400),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(), // Users have to be able to login, which may be done by anyone.
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).Annotations(
			entgql.RelayConnection(),
		),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
