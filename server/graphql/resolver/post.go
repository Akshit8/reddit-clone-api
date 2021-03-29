package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Akshit8/reddit-clone-api/pkg/entity"
	"github.com/Akshit8/reddit-clone-api/server/graphql/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePost) (*model.Post, error) {
	newPost := entity.Post{
		Title:       input.Title,
		Description: input.Description,
	}

	post, err := r.PostService.CreatePost(ctx, newPost)
	if err != nil {
		return nil, err
	}

	result := &model.Post{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}

	return result, nil
}

func (r *queryResolver) GetPostByID(ctx context.Context, id int) (*model.Post, error) {
	post, err := r.PostService.GetPostByID(ctx, id)
	if err != nil {
		return nil, err
	}

	result := &model.Post{
		ID:          post.ID,
		Title:       post.Title,
		Description: post.Description,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
	}

	return result, nil
}
