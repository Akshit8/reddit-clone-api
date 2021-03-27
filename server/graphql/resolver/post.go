package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Akshit8/reddit-clone-api/pkg/entity"
	"github.com/Akshit8/reddit-clone-api/server/graphql/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePost) (*entity.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePost(ctx context.Context, title string) (*entity.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePost(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPostByID(ctx context.Context, id string) (*entity.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPosts(ctx context.Context) ([]*entity.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
