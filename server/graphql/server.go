// Package graphql impls type-safe graphql server
package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Akshit8/reddit-clone-api/pkg/post"
	"github.com/Akshit8/reddit-clone-api/pkg/token"
	"github.com/Akshit8/reddit-clone-api/pkg/user"
	"github.com/Akshit8/reddit-clone-api/server/graphql/generated"
	"github.com/Akshit8/reddit-clone-api/server/graphql/middleware"
	"github.com/Akshit8/reddit-clone-api/server/graphql/resolver"
	chi "github.com/go-chi/chi/v5"
)

// NewGraphqlServer creates a new graphql server and returns the server multiplexer
func NewGraphqlServer(
	postService post.Service,
	userService user.Service,
	tokenMaker token.Maker,
) *chi.Mux {

	config := generated.Config{Resolvers: &resolver.Resolver{
		PostService: postService,
		UserService: userService,
	}}
	
	executableSchema := generated.NewExecutableSchema(config)
	srv := handler.NewDefaultServer(executableSchema)

	router := chi.NewRouter()

	router.Use(middleware.Auth(tokenMaker, userService))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	return router
}
