-- name: CreatePost :one
INSERT INTO posts (
  owner,
  title,
  content
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1 LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts WHERE created_at < $2 ORDER BY created_at DESC LIMIT $1;

-- name: GetAllUserPosts :many
SELECT * FROM posts WHERE owner = $1 ORDER BY id;

-- name: UpdatePostByID :exec
UPDATE posts SET title = $2, content = $3 WHERE id = $1;

-- name: DeletePostByID :exec
DELETE FROM posts WHERE id = $1;