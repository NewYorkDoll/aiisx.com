package resolver

import (
	"aiisx.com/src/database/graphql/graph/generated"
	"aiisx.com/src/ent"
	"aiisx.com/src/models"
	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the GQL resolver root.
type Resolver struct {
	codingStats *models.CodingStats
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{},
	})
}
