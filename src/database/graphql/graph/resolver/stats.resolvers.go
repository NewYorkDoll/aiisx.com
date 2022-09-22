package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"aiisx.com/src/models"
	"aiisx.com/src/wakapi"
)

// CodingStats is the resolver for the codingStats field.
func (r *queryResolver) CodingStats(ctx context.Context) (*models.CodingStats, error) {
	return wakapi.Statistics.Load(), nil
}
