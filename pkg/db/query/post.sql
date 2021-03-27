-- name: CreatePost :one
INSERT INTO posts (
  title,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3
) RETURNING *;