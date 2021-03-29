-- name: CreatePost :one
INSERT INTO posts (
  title,
  description
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetPostByID :one
SELECT * FROM posts WHERE id = $1 LIMIT 1;

-- -- name: GetPosts :many
-- SELECT * FROM posts ORDER BY id;

-- -- name: UpdatePostByID :exec
-- UPDATE posts SET title = $2 WHERE id = $1;

-- -- name: DeletePostByID :exec
-- DELETE FROM posts WHERE id = $1;