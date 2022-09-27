package graphql

import (
	"context"
	"errors"
	"strings"

	"aiisx.com/src/auth"
	"aiisx.com/src/ent"
	"aiisx.com/src/ent/privacy"
	"ariga.io/entcache"
	"github.com/99designs/gqlgen/graphql"
	"github.com/apex/log"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func injectClient(client *ent.Client) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		ctx = ent.NewContext(ctx, client)
		return next(ctx)
	}
}

func cacheEvictAdmin(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	if auth.RolesFromContext(ctx).Has("admin") {
		ctx = entcache.Evict(ctx)
	}

	return next(ctx)
}

func recoverFunc(ctx context.Context, err interface{}) error {
	logger := log.FromContext(ctx)

	switch e := err.(type) {
	case error:
		logger.WithError(e).Error("graphql panic")
	case string:
		logger.WithError(errors.New(e)).Error("graphql panic")
	default:
		logger.WithField("error", e).Error("graphql panic")
	}

	return gqlerror.Errorf("internal error")
}

func requestLogger(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	oc := graphql.GetOperationContext(ctx)

	log.FromContext(ctx).WithFields(log.Fields{
		"op":    oc.OperationName,
		"query": strings.ReplaceAll(strings.ReplaceAll(oc.RawQuery, "  ", ""), "\n", " "),
	}).Debug("handling request")

	return next(ctx)
}

func errorPresenter(ctx context.Context, err error) (gerr *gqlerror.Error) {
	defer func() {
		if errors.Is(err, privacy.Deny) {
			gerr.Message = "Permission denied"
		}
	}()

	if errors.As(err, &gerr) {
		if gerr.Path == nil {
			gerr.Path = graphql.GetPath(ctx)
		}

		return gerr
	}

	path := graphql.GetPath(ctx)
	logger := log.FromContext(ctx)

	logger.WithError(err).WithField("path", gerr.Path).Error("graphql internal failure")

	return gqlerror.ErrorPathf(path, "internal error")
}
