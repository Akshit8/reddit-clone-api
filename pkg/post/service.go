// Package post impls business logic for entity POST.
package post

import (
	"context"
	"time"

	db "github.com/Akshit8/reddit-clone-api/pkg/db/sqlc"
	"github.com/Akshit8/reddit-clone-api/pkg/entity"
)

// Service defines functions available on entity post
type Service interface {
	CreatePost(ctx context.Context, post entity.Post) (entity.Post, error)
	GetPostByID(ctx context.Context, id int) (entity.Post, error)
	GetPosts(ctx context.Context) ([]entity.Post, error)
	UpdatePostByID(ctx context.Context, post entity.Post) (entity.Post, error)
	DeletePostByID(ctx context.Context, id int) (bool, error)
}

type postService struct {
	repo *db.Queries
}

// NewPostService creates new instance of postService
func NewPostService(repo *db.Queries) Service {
	return &postService{
		repo: repo,
	}
}

func (p *postService) CreatePost(ctx context.Context, newPost entity.Post) (entity.Post, error) {
	createTimestamp := time.Now()
	
	createPostParams := db.CreatePostParams{
		Title: newPost.Title,
		Description: newPost.Description,
		CreatedAt: createTimestamp,
		UpdatedAt: createTimestamp,
	}

	post, err := p.repo.CreatePost(ctx, createPostParams)
	if err != nil {
		return entity.Post{}, nil
	}

	result := entity.Post{
		ID: int(post.ID),
		Title: post.Title,
		Description: post.Description,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	return result, nil
}

func (p *postService) GetPostByID(ctx context.Context, id int) (entity.Post, error) {
	post, err := p.repo.GetPostByID(ctx, int64(id))
	if err != nil {
		return entity.Post{}, err
	}

	result := entity.Post{
		ID: int(post.ID),
		Title: post.Title,
		Description: post.Description,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	return result, nil
}

func (p *postService) GetPosts(ctx context.Context) ([]entity.Post, error) {
	posts, err := p.repo.GetPosts(ctx)
	if err != nil {
		return []entity.Post{}, err
	}

	var result []entity.Post
	for _, post := range posts {
		result = append(result, entity.Post{
			ID: int(post.ID),
			Title: post.Title,
			Description: post.Description,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return result, nil
}

func updateHelper(post *entity.Post, updatedPost entity.Post) (title string, description string) {
	if updatedPost.Title != "" {
		title = updatedPost.Title
		post.Title = updatedPost.Title
	} else {
		title = post.Title
	}

	if updatedPost.Description != "" {
		description = updatedPost.Description
		post.Description = updatedPost.Description 
	} else {
		description = post.Description
	}

	return
}

func (p *postService) UpdatePostByID(ctx context.Context, updatedPost entity.Post) (entity.Post, error) {
	post, err := p.repo.GetPostByID(ctx, int64(updatedPost.ID))
	if err != nil {
		return entity.Post{}, err
	}
	
	newUpdatedTimeStamp := time.Now()

	result := entity.Post{
		ID: int(post.ID),
		Title: post.Title,
		Description: post.Description,
		CreatedAt: post.CreatedAt,
		UpdatedAt: newUpdatedTimeStamp,
	}

	title, description := updateHelper(&result, updatedPost)

	err = p.repo.UpdatePostByID(ctx, db.UpdatePostByIDParams{
		ID: int64(updatedPost.ID),
		Title: title,
		Description: description,
		UpdatedAt: newUpdatedTimeStamp,
	})
	if err != nil {
		return entity.Post{}, err
	}

	return result, nil
}

func (p *postService) DeletePostByID(ctx context.Context, id int) (bool, error) {
	_, err := p.repo.GetPostByID(ctx, int64(id))
	if err != nil {
		return false, err
	}

	err = p.repo.DeletePostByID(ctx, int64(id))
	if err != nil {
		return false, err
	}

	return true, err
}

