package post

import (
	"context"

	db "github.com/Akshit8/reddit-clone-api/pkg/db/sqlc"
	"github.com/Akshit8/reddit-clone-api/pkg/entity"
)

type Service interface {
	CreatePost(ctx context.Context, post entity.Post) (*db.Post, error)
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

func (p *postService) CreatePost(ctx context.Context, newPost entity.Post) (*db.Post, error) {
	post, err := p.repo.CreatePost(ctx, newPost.Title)
	if err != nil {
		return nil, err
	}
	return &post, nil
}
