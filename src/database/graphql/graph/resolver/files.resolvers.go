package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"aiisx.com/src/database/graphql/graph/generated"
	"aiisx.com/src/ent"
)

// CreateFiles is the resolver for the createFiles field.
func (r *mutationResolver) CreateFiles(ctx context.Context, input ent.CreateFilesInput) (*ent.Files, error) {
	panic(fmt.Errorf("not implemented: CreateFiles - createFiles"))
}

// UpdateFiles is the resolver for the updateFiles field.
func (r *mutationResolver) UpdateFiles(ctx context.Context, id int, input ent.UpdateFilesInput) (*ent.Files, error) {
	panic(fmt.Errorf("not implemented: UpdateFiles - updateFiles"))
}

// DeleteFiles is the resolver for the deleteFiles field.
func (r *mutationResolver) DeleteFiles(ctx context.Context, id int) (int, error) {
	panic(fmt.Errorf("not implemented: DeleteFiles - deleteFiles"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
