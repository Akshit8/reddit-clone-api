// Package post impls business logic for entity POST.
package post

import (
	"context"

	db "github.com/Akshit8/reddit-clone-api/pkg/db/sqlc"
	"github.com/Akshit8/reddit-clone-api/pkg/entity"
)

// Service defines functions available on entity post
type Service interface {
	CreatePost(ctx context.Context, post entity.Post) (entity.Post, error)
	GetPostByID(ctx context.Context, id int) (entity.Post, error)
	GetPosts(ctx context.Context) ([]entity.Post, error)
	// UpdatePost(ctx context.Context, id int, title string) (entity.Post, error)
	// DeletePost(ctx context.Context, id int) (bool, error)
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
	createPostParams := db.CreatePostParams{
		Title: newPost.Title,
		Description: newPost.Description,
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
	post, err := p.repo.GetPostByID(ctx, int32(id))
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

// func (p *postService) UpdatePost(ctx context.Context, id int, title string) (entity.Post, error) {
// 	err := p.repo.UpdatePostByID(ctx, db.UpdatePostByIDParams{
// 		ID: int32(id),
// 		Title: title,
// 	})
// 	if err != nil {
// 		return entity.Post{}, err
// 	}
// 	return entity.Post{}, nil
// }

// func (p *postService) DeletePost(ctx context.Context, id int) (bool, error) {
// 	err := p.repo.DeletePostByID(ctx, int32(id))
// 	if err != nil {
// 		return false, err
// 	}
// 	return true, err
// }

