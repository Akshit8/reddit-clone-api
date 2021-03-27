package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Akshit8/reddit-clone-api/server/graphql/generated"
)

func (r *mutationResolver) HealthMutation(ctx context.Context, msg string) (string, error) {
	return msg, nil
}

func (r *queryResolver) HealthQuery(ctx context.Context) (string, error) {
	return "health query working", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
