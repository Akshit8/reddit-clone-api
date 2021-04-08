package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Akshit8/reddit-clone-api/pkg/entity"
	"github.com/Akshit8/reddit-clone-api/server/graphql/generated"
	"github.com/Akshit8/reddit-clone-api/server/graphql/middleware"
	"github.com/Akshit8/reddit-clone-api/server/graphql/model"
	"github.com/Akshit8/reddit-clone-api/server/graphql/util"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePost) (*entity.Post, error) {
	user := middleware.FindUserFromContext(ctx)
	if user == nil {
		return nil, errors.New("user is unauthorized")
	}

	newPost := entity.Post{
		Owner:   user.ID,
		Title:   input.Title,
		Content: input.Content,
	}

	post, err := r.PostService.CreatePost(ctx, newPost)
	if err != nil {
		return nil, err
	}

	// result := &model.Post{
	// 	ID:        post.ID,
	// 	Title:     post.Title,
	// 	Content:   post.Content,
	// 	CreatedAt: post.CreatedAt,
	// 	UpdatedAt: post.UpdatedAt,
	// }

	return &post, nil
}

func (r *mutationResolver) UpdatePostByID(ctx context.Context, input model.UpdatePost) (*entity.Post, error) {
	if input.Title == nil && input.Content == nil {
		return nil, errors.New("no update field provided")
	}

	updatedPost := entity.Post{
		ID:      input.ID,
		Title:   util.StringPointerHelper(input.Title),
		Content: util.StringPointerHelper(input.Content),
	}
	fmt.Println(updatedPost)
	post, err := r.PostService.UpdatePostByID(ctx, updatedPost)
	if err != nil {
		return nil, err
	}

	// result := &model.Post{
	// 	ID:        post.ID,
	// 	Title:     post.Title,
	// 	Content:   post.Content,
	// 	CreatedAt: post.CreatedAt,
	// 	UpdatedAt: post.UpdatedAt,
	// }

	return &post, nil
}

func (r *mutationResolver) DeletePostByID(ctx context.Context, id int) (bool, error) {
	return r.PostService.DeletePostByID(ctx, id)
}

func (r *mutationResolver) UpvotePost(ctx context.Context, input model.UpvotePost) (bool, error) {
	user := middleware.FindUserFromContext(ctx)
	if user == nil {
		return false, errors.New("user is unauthorized")
	}
	return r.PostService.UpvotePost(ctx, input.ID, user.ID, input.Upvote)
}

func (r *postResolver) Owner(ctx context.Context, obj *entity.Post) (*entity.User, error) {
	log.Println("owner resolver: ", obj.Owner)
	log.Println("using")
	loggedInUser := middleware.FindUserFromContext(ctx)
	user, err := r.UserService.GetUserByID(ctx, obj.Owner)
	if loggedInUser.ID != user.ID {
		user.Email = ""
	}
	return &user, err
}

func (r *postResolver) ContentPreview(ctx context.Context, obj *entity.Post) (string, error) {
	if len(obj.Content) > 50 {
		return obj.Content[:50], nil
	}
	return obj.Content, nil
}

func (r *postResolver) VoteStatus(ctx context.Context, obj *entity.Post) (*int, error) {
	loggedInUser := middleware.FindUserFromContext(ctx)
	fmt.Println("owner:", loggedInUser.ID)
	upvote, err := r.UpvoteService.GetUpvote(ctx, loggedInUser.ID, obj.ID)
	if err != nil {
		return nil, nil
	}

	return &upvote.Value, nil
}

func (r *queryResolver) GetPostByID(ctx context.Context, id int) (*entity.Post, error) {
	post, err := r.PostService.GetPostByID(ctx, id)
	if err != nil {
		return nil, err
	}
	fmt.Println(post)
	return &post, nil
}

func (r *queryResolver) GetPosts(ctx context.Context, limit int, cursor *string) (*model.PaginatedPosts, error) {
	var arg time.Time
	var err error
	if cursor == nil {
		arg = time.Now()
		log.Println(arg)
	} else {
		arg, err = time.Parse("2006-01-02T15:04:05Z", *cursor)
		if err != nil {
			return nil, err
		}
	}
	posts, err := r.PostService.GetPosts(ctx, limit+1, arg)
	if err != nil {
		return nil, err
	}
	fmt.Println("post length", len(posts))
	resultLength := len(posts)
	if len(posts) > limit {
		resultLength--
	}
	result := make([]*entity.Post, resultLength)
	for i := range posts {
		fmt.Println(i)
		if i == limit {
			break
		}
		result[i] = &posts[i]
	}

	return &model.PaginatedPosts{
		Posts:   result,
		HasMore: limit+1 == len(posts),
	}, nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
