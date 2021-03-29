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

const getPostByID = `-- name: GetPostByID :one
SELECT id, title, description, created_at, updated_at FROM posts WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPostByID(ctx context.Context, id int32) (Post, error) {
	row := q.queryRow(ctx, q.getPostByIDStmt, getPostByID, id)
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

const getPosts = `-- name: GetPosts :many
SELECT id, title, description, created_at, updated_at FROM posts ORDER BY id
`

func (q *Queries) GetPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.query(ctx, q.getPostsStmt, getPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
