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
	GetPostByID(ctx context.Context, id string) (entity.Post, error)
	GetPosts(ctx context.Context) ([]entity.Post, error)
	UpdatePost(ctx context.Context, id string, title string) (entity.Post, error)
	DeletePost(ctx context.Context, id string) (bool, error)
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

func (p *postService) CreatePost(ctx context.Context, newPost db.Post) (db.Post, error) {
	post, err := p.repo.CreatePost(ctx, newPost.Title)
	if err != nil {
		return db.Post{}, err
	}
	return post, nil
}

func (p *postService) GetPostByID(ctx context.Context, id string) (db.Post, error) {
	return db.Post{}, nil
}

func (p *postService) GetPosts(ctx context.Context) ([]db.Post, error) {
	return []db.Post{}, nil
}
