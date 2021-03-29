package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

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

func (r *mutationResolver) UpdatePostByID(ctx context.Context, input model.UpdatePost) (*model.Post, error) {
	if input.Title == nil && input.Description == nil {
		return nil, errors.New("no update field provided")
	}

	updatedPost := entity.Post{
		ID:          input.ID,
		Title:       updateHelper(input.Title),
		Description: updateHelper(input.Description),
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func updateHelper(a *string) string {
	if a != nil {
		return *a
	}
	return ""
}
