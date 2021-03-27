// Package graphql impls type-safe graphql server
package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Akshit8/reddit-clone-api/pkg/post"
	"github.com/Akshit8/reddit-clone-api/server/graphql/generated"
	"github.com/Akshit8/reddit-clone-api/server/graphql/resolver"
)

// NewGraphqlServer creates a new graphql server and returns the server multiplexer
func NewGraphqlServer(postService post.Service) *http.ServeMux {
	config := generated.Config{Resolvers: &resolver.Resolver{
		PostService: postService,
	}}
	executableSchema := generated.NewExecutableSchema(config)
	srv := handler.NewDefaultServer(executableSchema)

	r := http.NewServeMux()

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	return r
}
