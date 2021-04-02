package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Akshit8/reddit-clone-api/pkg/entity"
	"github.com/Akshit8/reddit-clone-api/server/graphql/generated"
	"github.com/Akshit8/reddit-clone-api/server/graphql/model"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterUser) (*entity.User, error) {
	user, err := r.UserService.RegisterUser(ctx, input.Username, input.Password, input.Email)
	if err != nil {
		return nil, err
	}

	// result := &model.User{
	// 	ID:        user.ID,
	// 	Username:  user.Username,
	// 	Email:     user.Email,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// }

	return &user, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginUser) (*model.LoginResponse, error) {
	accessToken, err := r.UserService.LoginUser(ctx, input.UsernameOrEmail, input.Password)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{Token: accessToken}, nil
}

func (r *mutationResolver) ForgotPassword(ctx context.Context, email string) (bool, error) {
	return r.UserService.ForgotPassword(ctx, email)
}

func (r *mutationResolver) ChangePassword(ctx context.Context, input model.ChangePassword) (*model.LoginResponse, error) {
	token, err := r.UserService.ChangePassword(ctx, input.Token, input.NewPassword)
	if err != nil {
		return nil, err
	}
	return &model.LoginResponse{Token: token}, nil
}

func (r *queryResolver) Me(ctx context.Context, id int) (*entity.User, error) {
	user, err := r.UserService.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// result := &model.User{
	// 	ID:        user.ID,
	// 	Username:  user.Username,
	// 	Email:     user.Email,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// }

	return &user, nil
}

func (r *userResolver) Posts(ctx context.Context, obj *entity.User) ([]*entity.Post, error) {
	posts, err := r.PostService.GetUsersPost(ctx, obj.ID)

	var result []*entity.Post
	for _, post := range posts {
		result = append(result, &post)
	}

	return result, err
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
