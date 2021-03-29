// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username,
  password,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4
) RETURNING id, username, password, created_at, updated_at
`

type CreateUserParams struct {
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.Username,
		arg.Password,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, password, created_at, updated_at FROM users WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getUserByIDStmt, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, password, created_at, updated_at FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.queryRow(ctx, q.getUserByUsernameStmt, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
