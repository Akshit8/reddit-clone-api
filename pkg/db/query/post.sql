-- name: CreatePost :one
INSERT INTO posts (
  title,
  description,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1 LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts ORDER BY id;

-- name: GetAllUserPosts :many
SELECT * FROM posts WHERE owner = $1 ORDER BY id;

-- name: UpdatePostByID :exec
UPDATE posts SET title = $2, description = $3, updated_at = $4 WHERE id = $1;

-- name: DeletePostByID :exec
DELETE FROM posts WHERE id = $1;