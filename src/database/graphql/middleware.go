package graphql

import (
	"context"

	"aiisx.com/src/ent"
	"github.com/99designs/gqlgen/graphql"
)

func injectClient(client *ent.Client) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		ctx = ent.NewContext(ctx, client)
		return next(ctx)
	}
}
