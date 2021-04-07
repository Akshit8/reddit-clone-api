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
	GetPosts(ctx context.Context, limit int, cursor time.Time) ([]entity.Post, error)
	GetUsersPost(ctx context.Context, userID int) ([]entity.Post, error)
	UpdatePostByID(ctx context.Context, post entity.Post) (entity.Post, error)
	DeletePostByID(ctx context.Context, id int) (bool, error)
	UpvotePost(ctx context.Context, postID, userID int, upvote bool) (bool, error)
}

type postService struct {
	repo *db.Store
}

// NewPostService creates new instance of postService
func NewPostService(repo *db.Store) Service {
	return &postService{
		repo: repo,
	}
}

func (p *postService) CreatePost(ctx context.Context, newPost entity.Post) (entity.Post, error) {
	createPostParams := db.CreatePostParams{
		Owner:   int64(newPost.Owner),
		Title:   newPost.Title,
		Content: newPost.Content,
	}

	post, err := p.repo.CreatePost(ctx, createPostParams)
	if err != nil {
		return entity.Post{}, nil
	}

	result := entity.Post{
		ID:        int(post.ID),
		Title:     post.Title,
		Owner:     int(post.Owner),
		Content:   post.Content,
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
		ID:        int(post.ID),
		Title:     post.Title,
		Owner:     int(post.Owner),
		Content:   post.Content,
		UpVotes:   int(post.Upvotes),
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}

	return result, nil
}

func (p *postService) GetPosts(ctx context.Context, limit int, cursor time.Time) ([]entity.Post, error) {
	posts, err := p.repo.GetPosts(ctx, db.GetPostsParams{
		Limit:     int32(limit),
		CreatedAt: cursor,
	})
	if err != nil {
		return []entity.Post{}, err
	}

	var result []entity.Post
	for _, post := range posts {
		result = append(result, entity.Post{
			ID:        int(post.ID),
			Title:     post.Title,
			Owner:     int(post.Owner),
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return result, nil
}

func updateHelper(post *entity.Post, updatedPost entity.Post) (title string, content string) {
	if updatedPost.Title != "" {
		title = updatedPost.Title
		post.Title = updatedPost.Title
	} else {
		title = post.Title
	}

	if updatedPost.Content != "" {
		content = updatedPost.Content
		post.Content = updatedPost.Content
	} else {
		content = post.Content
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
		ID:        int(post.ID),
		Title:     post.Title,
		Owner:     int(post.Owner),
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
		UpdatedAt: newUpdatedTimeStamp,
	}

	title, content := updateHelper(&result, updatedPost)

	err = p.repo.UpdatePostByID(ctx, db.UpdatePostByIDParams{
		ID:      int64(updatedPost.ID),
		Title:   title,
		Content: content,
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

func (p *postService) GetUsersPost(ctx context.Context, userID int) ([]entity.Post, error) {
	posts, err := p.repo.GetAllUserPosts(ctx, int64(userID))
	if err != nil {
		return []entity.Post{}, err
	}

	var result []entity.Post
	for _, post := range posts {
		result = append(result, entity.Post{
			ID:        int(post.ID),
			Title:     post.Title,
			Owner:     int(post.Owner),
			Content:   post.Content,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		})
	}

	return result, nil
}

func (p *postService) UpvotePost(ctx context.Context, postID, userID int, upvote bool) (bool, error) {
	var value int
	if upvote {
		value = 1
	} else {
		value = -1
	}

	err := p.repo.UpvoteTx(ctx, db.UpvoteTxParams{
		UserID: int64(userID),
		PostID: int64(postID),
		Value:  int32(value),
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
