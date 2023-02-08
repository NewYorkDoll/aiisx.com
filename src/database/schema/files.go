package schema

import (
	"aiisx.com/src/ent/privacy"
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Files struct {
	ent.Schema
}

// 创建cos_files表
func (Files) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().Annotations(
			entgql.OrderField("NAME"),
		),
		field.String("url").Unique().Annotations(
			entgql.OrderField("URL"),
		),
		field.String("bucket").Unique().Annotations(
			entgql.OrderField("BUCKET"),
		),
	}
}

// 和posts表关联
func (Files) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).Annotations(
			entgql.RelayConnection(),
		),
	}
}

func (Files) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Files) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(), // Users have to be able to login, which may be done by anyone.
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}

func (Files) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
	}
}
