package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"aiisx.com/src/database/graphql/graph/generated"
	"aiisx.com/src/database/graphql/graph/model"
	"aiisx.com/src/ent"
)

// License is the resolver for the license field.
func (r *githubRepositoryResolver) License(ctx context.Context, obj *ent.GithubRepository) (*model.GithubLicense, error) {
	panic(fmt.Errorf("not implemented: License - license"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Githubevents is the resolver for the githubevents field.
func (r *queryResolver) Githubevents(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubEventOrder, where *ent.GithubEventWhereInput) (*ent.GithubEventConnection, error) {
	db := ent.FromContext(ctx)
	if db == nil {
		panic("database client is nil")
	}
	return db.GithubEvent.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithGithubEventOrder(orderBy),
		ent.WithGithubEventFilter(where.Filter),
	)
}

// Githubrepositories is the resolver for the githubrepositories field.
func (r *queryResolver) Githubrepositories(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.GithubRepositoryOrder, where *ent.GithubRepositoryWhereInput) (*ent.GithubRepositoryConnection, error) {
	panic(fmt.Errorf("not implemented: Githubrepositories - githubrepositories"))
}

// GithubRepository returns generated.GithubRepositoryResolver implementation.
func (r *Resolver) GithubRepository() generated.GithubRepositoryResolver {
	return &githubRepositoryResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type githubRepositoryResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

type githubEventResolver struct{ *Resolver }
