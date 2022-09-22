package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"aiisx.com/src/database/graphql/graph/generated"
	"aiisx.com/src/database/graphql/graph/model"
	"aiisx.com/src/gh"
	"github.com/google/go-github/v47/github"
)

// CreatedAt is the resolver for the createdAt field.
func (r *githubUserResolver) CreatedAt(ctx context.Context, obj *github.User) (*model.Timestamp, error) {
	panic(fmt.Errorf("not implemented: CreatedAt - createdAt"))
}

// UpdatedAt is the resolver for the updatedAt field.
func (r *githubUserResolver) UpdatedAt(ctx context.Context, obj *github.User) (*model.Timestamp, error) {
	panic(fmt.Errorf("not implemented: UpdatedAt - updatedAt"))
}

// GithubUser is the resolver for the githubUser field.
func (r *queryResolver) GithubUser(ctx context.Context) (*github.User, error) {
	user := gh.User.Load()
	if user == nil {
		return nil, fmt.Errorf("information not available yet")
	}

	return user, nil
}

// GithubUser returns generated.GithubUserResolver implementation.
func (r *Resolver) GithubUser() generated.GithubUserResolver { return &githubUserResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type githubUserResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
