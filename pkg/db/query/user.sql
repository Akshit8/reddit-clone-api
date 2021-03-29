-- name: CreateUser :one
INSERT INTO users (
  username,
  password,
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = $1 LIMIT 1;