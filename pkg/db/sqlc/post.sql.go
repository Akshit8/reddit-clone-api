// Code generated by sqlc. DO NOT EDIT.
// source: post.sql

package db

import (
	"context"
	"time"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
  owner,
  title,
  content
) VALUES (
  $1, $2, $3
) RETURNING id, owner, title, content, upvotes, created_at, updated_at
`

type CreatePostParams struct {
	Owner   int64
	Title   string
	Content string
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.queryRow(ctx, q.createPostStmt, createPost, arg.Owner, arg.Title, arg.Content)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Title,
		&i.Content,
		&i.Upvotes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePostByID = `-- name: DeletePostByID :exec
DELETE FROM posts WHERE id = $1
`

func (q *Queries) DeletePostByID(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deletePostByIDStmt, deletePostByID, id)
	return err
}

const getAllUserPosts = `-- name: GetAllUserPosts :many
SELECT id, owner, title, content, upvotes, created_at, updated_at FROM posts WHERE owner = $1 ORDER BY id
`

func (q *Queries) GetAllUserPosts(ctx context.Context, owner int64) ([]Post, error) {
	rows, err := q.query(ctx, q.getAllUserPostsStmt, getAllUserPosts, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Title,
			&i.Content,
			&i.Upvotes,
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

const getPostByID = `-- name: GetPostByID :one
SELECT id, owner, title, content, upvotes, created_at, updated_at FROM posts WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPostByID(ctx context.Context, id int64) (Post, error) {
	row := q.queryRow(ctx, q.getPostByIDStmt, getPostByID, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Title,
		&i.Content,
		&i.Upvotes,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :many
SELECT id, owner, title, content, upvotes, created_at, updated_at FROM posts WHERE created_at > $2 ORDER BY created_at DESC LIMIT $1
`

type GetPostsParams struct {
	Limit     int32
	CreatedAt time.Time
}

func (q *Queries) GetPosts(ctx context.Context, arg GetPostsParams) ([]Post, error) {
	rows, err := q.query(ctx, q.getPostsStmt, getPosts, arg.Limit, arg.CreatedAt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Title,
			&i.Content,
			&i.Upvotes,
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

const updatePostByID = `-- name: UpdatePostByID :exec
UPDATE posts SET title = $2, content = $3 WHERE id = $1
`

type UpdatePostByIDParams struct {
	ID      int64
	Title   string
	Content string
}

func (q *Queries) UpdatePostByID(ctx context.Context, arg UpdatePostByIDParams) error {
	_, err := q.exec(ctx, q.updatePostByIDStmt, updatePostByID, arg.ID, arg.Title, arg.Content)
	return err
}
