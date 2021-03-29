package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Akshit8/reddit-clone-api/server/graphql/model"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterUser) (*model.User, error) {
	user, err := r.UserService.RegisterUser(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	result := &model.User{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return result, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginUser) (*model.LoginResponse, error) {
	accessToken, err := r.UserService.LoginUser(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	result := &model.LoginResponse{
		Token: accessToken,
	}

	return result, nil
}
