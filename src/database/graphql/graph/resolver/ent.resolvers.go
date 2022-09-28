package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"aiisx.com/src/database/graphql/graph/generated"
	"aiisx.com/src/database/graphql/graph/model"
	"aiisx.com/src/ent"
)

// License is the resolver for the license field.
func (r *githubRepositoryResolver) License(ctx context.Context, obj *ent.GithubRepository) (*model.GithubLicense, error) {
	panic(fmt.Errorf("not implemented: License - license"))
}

// CreateTime is the resolver for the createTime field.
func (r *labelResolver) CreateTime(ctx context.Context, obj *ent.Label) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: CreateTime - createTime"))
}

// UpdateTime is the resolver for the updateTime field.
func (r *labelResolver) UpdateTime(ctx context.Context, obj *ent.Label) (*time.Time, error) {
	panic(fmt.Errorf("not implemented: UpdateTime - updateTime"))
}

// Name is the resolver for the name field.
func (r *labelResolver) Name(ctx context.Context, obj *ent.Label) (string, error) {
	panic(fmt.Errorf("not implemented: Name - name"))
}

// Ping is the resolver for the ping field.
func (r *mutationResolver) Ping(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Ping - ping"))
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

// Labels is the resolver for the labels field.
func (r *queryResolver) Labels(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, where *ent.LabelWhereInput) (*ent.LabelConnection, error) {
	panic(fmt.Errorf("not implemented: Labels - labels"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// CreateTime is the resolver for the createTime field.
func (r *createLabelInputResolver) CreateTime(ctx context.Context, obj *ent.CreateLabelInput, data *time.Time) error {
	panic(fmt.Errorf("not implemented: CreateTime - createTime"))
}

// UpdateTime is the resolver for the updateTime field.
func (r *createLabelInputResolver) UpdateTime(ctx context.Context, obj *ent.CreateLabelInput, data *time.Time) error {
	panic(fmt.Errorf("not implemented: UpdateTime - updateTime"))
}

// Name is the resolver for the name field.
func (r *createLabelInputResolver) Name(ctx context.Context, obj *ent.CreateLabelInput, data string) error {
	panic(fmt.Errorf("not implemented: Name - name"))
}

// GithubRepositoryIDs is the resolver for the githubRepositoryIDs field.
func (r *createLabelInputResolver) GithubRepositoryIDs(ctx context.Context, obj *ent.CreateLabelInput, data []string) error {
	panic(fmt.Errorf("not implemented: GithubRepositoryIDs - githubRepositoryIDs"))
}

// GithubRepository returns generated.GithubRepositoryResolver implementation.
func (r *Resolver) GithubRepository() generated.GithubRepositoryResolver {
	return &githubRepositoryResolver{r}
}

// Label returns generated.LabelResolver implementation.
func (r *Resolver) Label() generated.LabelResolver { return &labelResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// CreateLabelInput returns generated.CreateLabelInputResolver implementation.
func (r *Resolver) CreateLabelInput() generated.CreateLabelInputResolver {
	return &createLabelInputResolver{r}
}

type githubRepositoryResolver struct{ *Resolver }
type labelResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type createLabelInputResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
type githubEventResolver struct{ *Resolver }
