-- name: CreatePost :one
INSERT INTO post (
  title
) VALUES (
  $1
) RETURNING *;