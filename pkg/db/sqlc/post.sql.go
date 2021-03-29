// Code generated by sqlc. DO NOT EDIT.
// source: post.sql

package db

import (
	"context"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
  title,
  description
) VALUES (
  $1, $2
) RETURNING id, title, description, created_at, updated_at
`

type CreatePostParams struct {
	Title       string
	Description string
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.queryRow(ctx, q.createPostStmt, createPost, arg.Title, arg.Description)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
