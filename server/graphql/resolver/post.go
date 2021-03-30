package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/Akshit8/reddit-clone-api/pkg/entity"
	"github.com/Akshit8/reddit-clone-api/pkg/middleware"
	"github.com/Akshit8/reddit-clone-api/server/graphql/model"
	"github.com/Akshit8/reddit-clone-api/server/graphql/util"
)

// ErrUserUnauthorized is ...
var ErrUserUnauthorized = errors.New("user is unauthorized")

func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePost) (*model.Post, error) {
	user := middleware.FindUserFromContext(ctx)
	if user == nil {
		return nil, ErrUserUnauthorized
	}

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

func (r *mutationResolver) UpdatePostByID(ctx context.Context, input model.UpdatePost) (*model.Post, error) {
	if input.Title == nil && input.Description == nil {
		return nil, errors.New("no update field provided")
	}

	updatedPost := entity.Post{
		ID:          input.ID,
		Title:       util.StringPointerHelper(input.Title),
		Description: util.StringPointerHelper(input.Description),
	}
	fmt.Println(updatedPost)
	post, err := r.PostService.UpdatePostByID(ctx, updatedPost)
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

func (r *mutationResolver) DeletePostByID(ctx context.Context, id int) (bool, error) {
	return r.PostService.DeletePostByID(ctx, id)
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

func (r *queryResolver) GetPosts(ctx context.Context) ([]*model.Post, error) {
	posts, err := r.PostService.GetPosts(ctx)
	if err != nil {
		return nil, err
	}

	var result []*model.Post
	for _, post := range posts {
		result = append(result, &model.Post{
			ID:          post.ID,
			Title:       post.Title,
			Description: post.Description,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
		})
	}

	return result, nil
}
