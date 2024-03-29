package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"aiisx.com/src/database"
	"aiisx.com/src/database/graphql/graph/generated"
	"aiisx.com/src/database/graphql/graph/model"
	"aiisx.com/src/ent"
	"ariga.io/entcache"
)

// License is the resolver for the license field.
func (r *githubRepositoryResolver) License(ctx context.Context, obj *ent.GithubRepository) (*model.GithubLicense, error) {
	panic(fmt.Errorf("not implemented: License - license"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	return ent.FromContext(ctx).Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	return ent.FromContext(ctx).Noders(ctx, ids)
}

// Filesslice is the resolver for the filesslice field.
func (r *queryResolver) Filesslice(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.FilesOrder, where *ent.FilesWhereInput) (*ent.FilesConnection, error) {
	panic(fmt.Errorf("not implemented: Filesslice - filesslice"))
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
func (r *queryResolver) Labels(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.LabelOrder, where *ent.LabelWhereInput) (*ent.LabelConnection, error) {
	db := ent.FromContext(ctx)
	labelQ := db.Label.Query()
	labels, err := labelQ.Paginate(
		// entcache.Skip(ctx) skip cache
		entcache.Skip(ctx), after, first, before, last,
		ent.WithLabelFilter(where.Filter),
	)
	return labels, err
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.PostOrder, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	conn, err := ent.FromContext(ctx).Post.Query().Paginate(
		ctx, after, first, before, last,
		ent.WithPostOrder(orderBy),
		ent.WithPostFilter(where.Filter),
	)

	if err == nil && conn.TotalCount == 1 {
		post := conn.Edges[0].Node

		if post.ContentHTML == "" {
			return conn, nil
		}

		go database.PostViewCounter(ctx, post)
	}
	return conn, err
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	panic(fmt.Errorf("not implemented: Users - users"))
}

// GithubRepository returns generated.GithubRepositoryResolver implementation.
func (r *Resolver) GithubRepository() generated.GithubRepositoryResolver {
	return &githubRepositoryResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type githubRepositoryResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
