package graphql

import (
	"aiisx.com/src/database/graphql/graph/resolver"
	"aiisx.com/src/ent"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

func New(db *ent.Client) *handler.Server {

	srv := handler.NewDefaultServer(resolver.NewSchema(db))

	srv.AroundOperations(requestLogger)
	srv.SetRecoverFunc(recoverFunc)

	srv.AroundOperations(injectClient(db))
	srv.AroundOperations(cacheEvictAdmin)

	srv.SetErrorPresenter(errorPresenter)

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))
	return srv
}
