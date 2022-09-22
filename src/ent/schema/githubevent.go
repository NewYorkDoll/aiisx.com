package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/go-github/v47/github"
)

// GithubEvent holds the schema definition for the GithubEvent entity.
type GithubEvent struct {
	ent.Schema
}

// Fields of the GithubEvent.
func (GithubEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("event_id").Unique().Annotations(
			entgql.OrderField("EVENT_ID"),
		),
		field.String("event_type").NotEmpty().Annotations(
			entgql.OrderField("EVENT_TYPE"),
		),
		field.Time("created_at").Annotations(
			entgql.OrderField("CREATED_AT"),
		),
		field.Bool("public").Default(false),
		field.Int64("actor_id").Annotations(
			entgql.OrderField("ACTOR_ID"),
		),
		field.JSON("actor", &github.User{}).Annotations(entgql.Type("GithubUser")),
		field.Int64("repo_id").Positive().Annotations(
			entgql.OrderField("REPO_ID"),
		),
		field.JSON("repo", &github.Repository{}).Annotations(entgql.Type("GithubEventRepo")),
		field.JSON("payload", map[string]any{}).Annotations(entgql.Type("Map")),
	}
}
