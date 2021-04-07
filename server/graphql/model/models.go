// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/Akshit8/reddit-clone-api/pkg/entity"
)

type ChangePassword struct {
	Token       string `json:"token"`
	NewPassword string `json:"newPassword"`
}

type CreatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginUser struct {
	UsernameOrEmail string `json:"usernameOrEmail"`
	Password        string `json:"password"`
}

type PaginatedPosts struct {
	Posts   []*entity.Post `json:"posts"`
	HasMore bool           `json:"hasMore"`
}

type RegisterUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdatePost struct {
	ID      int     `json:"id"`
	Title   *string `json:"title"`
	Content *string `json:"content"`
}

type UpvotePost struct {
	ID     int  `json:"id"`
	Upvote bool `json:"upvote"`
}
