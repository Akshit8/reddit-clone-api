-- name: CreatePost :one
INSERT INTO post (
  title
) VALUES (
  $1
) RETURNING *;

-- name: GetPostByID :one
SELECT * FROM post
WHERE id = $1 LIMIT 1;

-- name: GetPosts :many
SELECT * FROM post
ORDER BY id;