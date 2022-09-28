package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"aiisx.com/src/auth"
	"aiisx.com/src/ent"
)

// Self is the resolver for the self field.
func (r *queryResolver) Self(ctx context.Context) (*ent.User, error) {
	user := auth.User.Load()
	if user == nil {
		return nil, nil
	}

	return user, nil
}
