// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package resolver

//go:generate go run github.com/99designs/gqlgen generate
import (
	"aiisx.com/src/database/graphql/graph/generated"
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
func NewSchema() graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{},
	})
}
