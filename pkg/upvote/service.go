// Package upvote exposes functionality related to upvotes
package upvote

import (
	"context"

	db "github.com/Akshit8/reddit-clone-api/pkg/db/sqlc"
	"github.com/Akshit8/reddit-clone-api/pkg/entity"
)

// Service defines functions available on entity upvote
type Service interface {
	GetUpvote(ctx context.Context, userID, postID int) (*entity.Upvote, error)
}

type upvoteService struct {
	repo *db.Store
}

// NewUpvoteService creates new instance of upvoteService
func NewUpvoteService(repo *db.Store) Service {
	return &upvoteService{
		repo: repo,
	}
}

func (u *upvoteService) GetUpvote(ctx context.Context, userID, postID int) (*entity.Upvote, error) {
	upvote, err := u.repo.GetUpvote(ctx, db.GetUpvoteParams{UserId: int64(userID), PostId: int64(postID)})
	if err != nil {
		return nil, err;
	}

	result := &entity.Upvote{
		UserID: int(upvote.UserId),
		PostID: int(upvote.PostId),
		Value: int(upvote.Value),
		CreatedAt: upvote.CreatedAt,
		UpdatedAt: upvote.UpdatedAt,
	}

	return result, nil
}
